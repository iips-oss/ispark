package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/iips-oss/ispark/api/config"
	"github.com/iips-oss/ispark/api/models"
	"github.com/iips-oss/ispark/api/utils"
	"gorm.io/gorm"
)

func errJSON(c *fiber.Ctx, status int, msg string) error {
	return c.Status(status).JSON(fiber.Map{"error": msg})
}

func AdminLogin(c *fiber.Ctx) error {
	var input models.AdminLoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request body"})
	}
	if input.AdminID == "" || input.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Admin ID and Password are required"})
	}
	var admin models.Admin
	if err := config.DB.Where("admin_id = ?", input.AdminID).First(&admin).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}
	if !utils.CheckPasswordHash(input.Password, admin.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}
	accessToken, err := utils.GenerateAccessToken(admin.AdminID, admin.Email, admin.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate access token"})
	}
	return c.JSON(fiber.Map{
		"message":              "Admin logged in successfully",
		"access_token":         accessToken,
		"must_change_password": admin.MustChangePassword,
		"admin": fiber.Map{
			"admin_id": admin.AdminID,
			"name":     admin.Name,
			"role":     admin.Role,
		},
	})
}

// 1. POST /admin/change-password
func AdminChangePassword(c *fiber.Ctx) error {
	var input models.ChangePasswordInput
	if err := c.BodyParser(&input); err != nil {
		return errJSON(c, fiber.StatusBadRequest, "Cannot parse request body")
	}
	if input.CurrentPassword == "" || input.NewPassword == "" || input.ConfirmPassword == "" {
		return errJSON(c, fiber.StatusBadRequest, "All fields are required")
	}
	if input.NewPassword != input.ConfirmPassword {
		return errJSON(c, fiber.StatusBadRequest, "Passwords do not match")
	}
	adminID, ok := c.Locals("roll_no").(string)
	if !ok || adminID == "" {
		return errJSON(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	var admin models.Admin
	if err := config.DB.Where("admin_id = ?", adminID).First(&admin).Error; err != nil {
		return errJSON(c, fiber.StatusNotFound, "Admin not found")
	}

	if !utils.CheckPasswordHash(input.CurrentPassword, admin.Password) {
		return errJSON(c, fiber.StatusUnauthorized, "Current password is incorrect")
	}

	newHash, err := utils.HashPassword(input.NewPassword)
	if err != nil {
		return errJSON(c, fiber.StatusInternalServerError, "Failed to hash password")
	}

	admin.Password = newHash
	admin.MustChangePassword = false

	if err := config.DB.Save(&admin).Error; err != nil {
		return errJSON(c, fiber.StatusInternalServerError, "Failed to update password")
	}

	return c.JSON(fiber.Map{"message": "Password changed successfully"})
}

func getAuthenticatedAdmin(c *fiber.Ctx) (*models.Admin, error) {
	adminID, ok := c.Locals("roll_no").(string)
	if !ok || adminID == "" {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	var admin models.Admin
	if err := config.DB.Where("admin_id = ?", adminID).First(&admin).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	return &admin, nil
}

func scopeToAssignedBatch(query *gorm.DB, admin *models.Admin) (*gorm.DB, bool) {
	if admin.Role != "admin" {
		return query, true
	}
	if admin.AssignedBatch == "" {
		return query, false
	}

	return query.Where("roll_no LIKE ?", admin.AssignedBatch+"%"), true
}

func applyStudentStats(student *models.Student) {
	credits := 0
	pending := 0
	for _, cert := range student.Certificates {
		if cert.Status == "Approved" {
			credits += cert.Credits
		}
		if cert.Status == "Pending" {
			pending++
		}
	}

	student.CreditsEarned = credits
	student.PendingCertificates = pending
	student.ActivityCount = len(student.Enrollments)

	switch {
	case len(student.Enrollments) == 0 && len(student.Certificates) == 0:
		student.EngagementStatus = "Inactive"
	case pending > 0:
		student.EngagementStatus = "Pending Review"
	default:
		student.EngagementStatus = "Active"
	}
}

// 2. GET /api/admin/students -> View assigned students
func GetAllStudents(c *fiber.Ctx) error {
	currentUser, err := getAuthenticatedAdmin(c)
	if err != nil {
		return err
	}

	type StudentRow struct {
		models.Student
		CreditsEarned       int
		PendingCertificates int
		TotalCertificates   int
		ActivityCount       int
	}

	dbQuery := config.DB.Model(&models.Student{}).Select(`
		students.*,
		(SELECT COALESCE(SUM(credits), 0) FROM certificates WHERE certificates.student_roll_no = students.roll_no AND status = 'Approved') as credits_earned,
		(SELECT COUNT(*) FROM certificates WHERE certificates.student_roll_no = students.roll_no AND status = 'Pending') as pending_certificates,
		(SELECT COUNT(*) FROM certificates WHERE certificates.student_roll_no = students.roll_no) as total_certificates,
		(SELECT COUNT(*) FROM enrollments WHERE enrollments.student_roll_no = students.roll_no) as activity_count
	`)

	dbQuery, scoped := scopeToAssignedBatch(dbQuery, currentUser)
	if !scoped {
		return c.JSON(fiber.Map{"students": []models.Student{}})
	}

	var rows []StudentRow
	if err := dbQuery.Find(&rows).Error; err != nil {
		return errJSON(c, fiber.StatusInternalServerError, "Failed to retrieve students")
	}

	students := make([]models.Student, len(rows))
	for i, r := range rows {
		students[i] = r.Student
		students[i].CreditsEarned = r.CreditsEarned
		students[i].PendingCertificates = r.PendingCertificates
		students[i].TotalCertificates = r.TotalCertificates
		students[i].ActivityCount = r.ActivityCount

		if students[i].ActivityCount == 0 && students[i].CreditsEarned == 0 && students[i].PendingCertificates == 0 {
			students[i].EngagementStatus = "Inactive"
		} else if students[i].PendingCertificates > 0 {
			students[i].EngagementStatus = "Pending Review"
		} else {
			students[i].EngagementStatus = "Active"
		}
	}

	return c.JSON(fiber.Map{"students": students})
}

// 3. GET /api/admin/students/:roll -> One student's detail
func GetStudentDetail(c *fiber.Ctx) error {
	roll := c.Params("roll")

	currentUser, err := getAuthenticatedAdmin(c)
	if err != nil {
		return err
	}

	query, scoped := scopeToAssignedBatch(
		config.DB.Preload("Enrollments.Activity").
			Preload("Certificates").
			Where("roll_no = ?", roll),
		currentUser,
	)
	if !scoped {
		return errJSON(c, fiber.StatusNotFound, "Student not found")
	}

	var student models.Student
	if err := query.First(&student).Error; err != nil {
		return errJSON(c, fiber.StatusNotFound, "Student not found")
	}
	applyStudentStats(&student)

	return c.JSON(fiber.Map{"student": student})
}

// Fetches student if they exist AND are within admin's batch scope
func getScopedStudent(c *fiber.Ctx, roll string, admin *models.Admin) (*models.Student, error) {
	var student models.Student
	dbQuery, scoped := scopeToAssignedBatch(config.DB.Where("roll_no = ?", roll), admin)

	if !scoped || dbQuery.First(&student).Error != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Student not found or unauthorized")
	}
	return &student, nil
}

func GetMentorObservations(c *fiber.Ctx) error {
	roll := c.Params("roll")
	admin, err := getAuthenticatedAdmin(c)
	if err != nil {
		return err
	}

	if _, err := getScopedStudent(c, roll, admin); err != nil {
		return errJSON(c, fiber.StatusNotFound, err.Error())
	}

	var observations []models.AdminNote
	config.DB.Where("student_roll_no = ?", roll).Order("created_at asc").Find(&observations)
	return c.JSON(fiber.Map{"observations": observations})
}

func AddMentorObservation(c *fiber.Ctx) error {
	roll := c.Params("roll")
	admin, err := getAuthenticatedAdmin(c)
	if err != nil {
		return err
	}

	if _, err := getScopedStudent(c, roll, admin); err != nil {
		return errJSON(c, fiber.StatusNotFound, err.Error())
	}

	var input models.ObservationInput
	if err := c.BodyParser(&input); err != nil || input.Text == "" {
		return errJSON(c, fiber.StatusBadRequest, "Observation text is required")
	}

	note := models.AdminNote{
		StudentRollNo: roll,
		AdminID:       admin.AdminID,
		AuthorName:    admin.Name,
		Role:          "Mentor",
		Text:          input.Text,
	}

	if err := config.DB.Create(&note).Error; err != nil {
		return errJSON(c, fiber.StatusInternalServerError, "Failed to save observation")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Observation added", "observation": note})
}

func EditMentorObservation(c *fiber.Ctx) error {
	roll := c.Params("roll")
	id := c.Params("id")
	admin, err := getAuthenticatedAdmin(c)
	if err != nil {
		return err
	}

	if _, err := getScopedStudent(c, roll, admin); err != nil {
		return errJSON(c, fiber.StatusNotFound, err.Error())
	}

	var input models.ObservationInput
	if err := c.BodyParser(&input); err != nil || input.Text == "" {
		return errJSON(c, fiber.StatusBadRequest, "Observation text is required")
	}

	var note models.AdminNote
	if err := config.DB.Where("id = ? AND student_roll_no = ?", id, roll).First(&note).Error; err != nil {
		return errJSON(c, fiber.StatusNotFound, "Observation not found")
	}

	if note.AdminID != admin.AdminID {
		return errJSON(c, fiber.StatusForbidden, "Cannot edit another admin's observation")
	}

	note.Text = input.Text
	if err := config.DB.Save(&note).Error; err != nil {
		return errJSON(c, fiber.StatusInternalServerError, "Failed to update observation")
	}

	return c.JSON(fiber.Map{"message": "Observation updated", "observation": note})
}

func SendStudentNotice(c *fiber.Ctx) error {
	roll := c.Params("roll")
	admin, err := getAuthenticatedAdmin(c)
	if err != nil {
		return err
	}

	student, err := getScopedStudent(c, roll, admin)
	if err != nil {
		return errJSON(c, fiber.StatusNotFound, err.Error())
	}

	msg := fmt.Sprintf("Dear %s,\n\nOfficial notice regarding your iSPARC progress.\n\nRegards,\n%s", student.Name, admin.Name)

	if err := utils.SendEmail(student.EmailID, "iSPARC Notice", msg); err != nil {
		return errJSON(c, fiber.StatusInternalServerError, "Failed to dispatch email")
	}

	return c.JSON(fiber.Map{"message": "Notice sent successfully"})
}

func GetAdminDashboardStats(c *fiber.Ctx) error {
	admin, err := getAuthenticatedAdmin(c)
	if err != nil {
		return err
	}

	var totalStudents int64
	query, scoped := scopeToAssignedBatch(config.DB.Model(&models.Student{}), admin)
	if !scoped {
		return c.JSON(fiber.Map{
			"total_students":  0,
			"active_students": 0,
			"pending_reviews": 0,
			"average_credits": 0,
		})
	}
	query.Count(&totalStudents)

	// Pending reviews (Certificates with Status = 'Pending' for this batch)
	var pendingReviews int64
	certQuery := config.DB.Model(&models.Certificate{}).Joins("JOIN students on students.roll_no = certificates.student_roll_no").Where("certificates.status = ?", "Pending")
	if admin.Role == "admin" && admin.AssignedBatch != "" {
		certQuery = certQuery.Where("students.roll_no LIKE ?", admin.AssignedBatch+"%")
	}
	certQuery.Count(&pendingReviews)

	// Total Credits Earned
	var totalCredits sql.NullFloat64
	creditsQuery := config.DB.Model(&models.Certificate{}).Select("SUM(certificates.credits)").Joins("JOIN students on students.roll_no = certificates.student_roll_no").Where("certificates.status = ?", "Approved")
	if admin.Role == "admin" && admin.AssignedBatch != "" {
		creditsQuery = creditsQuery.Where("students.roll_no LIKE ?", admin.AssignedBatch+"%")
	}
	creditsQuery.Scan(&totalCredits)

	avgCredits := 0.0
	if totalStudents > 0 && totalCredits.Valid {
		avgCredits = totalCredits.Float64 / float64(totalStudents)
	}

	// Active Students
	var activeStudents int64
	activeQuery := config.DB.Model(&models.Student{}).
		Where("EXISTS (SELECT 1 FROM enrollments WHERE enrollments.student_roll_no = students.roll_no) OR EXISTS (SELECT 1 FROM certificates WHERE certificates.student_roll_no = students.roll_no)")
	if admin.Role == "admin" && admin.AssignedBatch != "" {
		activeQuery = activeQuery.Where("students.roll_no LIKE ?", admin.AssignedBatch+"%")
	}
	activeQuery.Count(&activeStudents)

	return c.JSON(fiber.Map{
		"total_students":  totalStudents,
		"active_students": activeStudents,
		"pending_reviews": pendingReviews,
		"average_credits": avgCredits,
	})
}

func GetRecentActivities(c *fiber.Ctx) error {
	admin, err := getAuthenticatedAdmin(c)
	if err != nil {
		return err
	}

	if admin.Role == "admin" && admin.AssignedBatch == "" {
		return c.JSON(fiber.Map{"recent_activities": []models.Certificate{}})
	}

	var recentCerts []models.Certificate
	certQuery := config.DB.Preload("Student").
		Joins("JOIN students on students.roll_no = certificates.student_roll_no").
		Order("certificates.created_at desc").
		Limit(5)

	if admin.Role == "admin" && admin.AssignedBatch != "" {
		certQuery = certQuery.Where("students.roll_no LIKE ?", admin.AssignedBatch+"%")
	}

	if err := certQuery.Find(&recentCerts).Error; err != nil {
		return errJSON(c, fiber.StatusInternalServerError, "Failed to fetch recent activities")
	}

	return c.JSON(fiber.Map{
		"recent_activities": recentCerts,
	})
}
