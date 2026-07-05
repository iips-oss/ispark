package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iips-oss/ispark/api/config"
	"github.com/iips-oss/ispark/api/models"
	"github.com/iips-oss/ispark/api/utils"
)

func AdminLogin(c *fiber.Ctx) error {
	var input models.AdminLoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request body"})
	}
	if input.AdminID == "" || input.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Admin ID and Password are required"})
	}
	// Find the admin using your existing Admin model
	var admin models.Admin
	if err := config.DB.Where("admin_id = ?", input.AdminID).First(&admin).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}
	// Verify Password
	if !utils.CheckPasswordHash(input.Password, admin.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}
	// Generate Access Token
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

// 1. POST /admin/change-password -> ChangePassword lets a logged-in admin set a new password (used for both voluntary changes and the forced first-login reset flow)
func ChangePassword(c *fiber.Ctx) error {
	var input models.ChangePasswordInput
	if err := c.BodyParser(&input); err != nil || input.NewPassword != input.ConfirmPassword {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input or password mismatch"})
	}

	var admin models.Admin
	adminID := c.Locals("roll_no").(string)
	if err := config.DB.First(&admin, "admin_id = ?", adminID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Account not found"})
	}

	if !utils.CheckPasswordHash(input.CurrentPassword, admin.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Incorrect current password"})
	}

	newHash, _ := utils.HashPassword(input.NewPassword)
	config.DB.Model(&admin).Update("password", newHash)

	return c.JSON(fiber.Map{"message": "Password changed successfully"})
}

// 2. GET /api/admin/students -> View assigned students
func GetAllStudents(c *fiber.Ctx) error {
	var students []models.Student
	if err := config.DB.Select("roll_no, name, course_name, semester, email_id").Find(&students).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve students"})
	}
	return c.JSON(fiber.Map{"students": students})
}

// 3. GET /api/admin/students/:roll -> One student's detail
func GetStudentDetail(c *fiber.Ctx) error {
	rollNo := c.Params("roll")
	var student models.Student

	if err := config.DB.Where("roll_no = ?", rollNo).First(&student).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Student not found"})
	}
	return c.JSON(fiber.Map{"student": student})
}
