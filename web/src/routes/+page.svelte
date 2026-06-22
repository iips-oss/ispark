<script lang="ts">
	import Navbar from '$lib/components/Navbar.svelte';
</script>

<svelte:head>
	<title>ISPARC | IIPS DAVV Cell</title>
	<meta name="description" content="Official portal of IIPS Skill, Personality Advancement & Refinement Cell (ISPARC). Facilitating student skill development and extracurricular grading." />
</svelte:head>

<div class="min-h-screen bg-academic-cream text-[#0f172a] flex flex-col font-sans">
	<!-- Navbar Mount -->
	<Navbar />

	<!-- Main Page Container -->
	<main class="flex-grow max-w-7xl w-full mx-auto px-6 sm:px-8 py-16 md:py-24 space-y-20">
		
		<!-- Editorial Two-Column Hero -->
		<section class="grid grid-cols-1 lg:grid-cols-12 gap-12 items-center" id="overview">
			
			<!-- Left text column (col span 7) -->
			<div class="lg:col-span-7 space-y-8">
				
				<!-- Academic badge -->
				<div class="inline-flex px-3 py-1.5 border border-slate-300 text-xs font-bold uppercase tracking-wider text-slate-600 bg-transparent rounded-none font-sans">
					Departmental Affiliated Cell
				</div>

				<h1 class="text-4xl sm:text-5xl lg:text-6xl font-normal leading-tight text-[#0f172a] text-serif">
					Skill, Personality <br />
					Advancement & <br />
					Refinement Cell.
				</h1>

				<p class="max-w-xl text-sm sm:text-base text-slate-600 leading-relaxed font-sans font-medium">
					Welcome to ISPARC, an official cell of the Institute of Infrastructure and Professional Studies (IIPS), DAVV. We facilitate structured extracurricular activities, personality enhancement tracks, and grading systems to foster student growth and leadership.
				</p>

				<!-- Call to actions -->
				<div class="flex flex-wrap gap-4 pt-2">
					<a
						href="#grading"
						class="px-6 py-3 text-xs font-bold btn-academic-primary rounded-none font-sans uppercase tracking-wider"
					>
						Grading Matrix
					</a>
					<a
						href="#tracks"
						class="px-6 py-3 text-xs font-bold btn-academic-secondary rounded-none font-sans uppercase tracking-wider"
					>
						Explore Tracks
					</a>
				</div>
			</div>

			<!-- Right stats/matrix column (col span 5) -->
			<div class="lg:col-span-5 bg-white border border-[#cbd5e1] p-8 shadow-xs rounded-none space-y-6">
				<h3 class="text-lg font-bold text-[#0f172a] text-serif border-b border-[#cbd5e1] pb-3">
					Activity Grading Matrix
				</h3>
				<p class="text-xs text-slate-500 font-medium font-sans">
					Extracurricular points are mapped directly to grades in semester results under the DAVV guidelines.
				</p>

				<!-- Matrix table -->
				<div class="space-y-3.5 pt-2 font-sans">
					<div class="flex justify-between items-center text-xs">
						<span class="font-bold text-slate-700">60+ Points</span>
						<span class="px-2.5 py-1 bg-emerald-50 text-emerald-800 font-bold border border-emerald-200">Grade O (Outstanding)</span>
					</div>
					<div class="flex justify-between items-center text-xs">
						<span class="font-bold text-slate-700">45 – 59 Points</span>
						<span class="px-2.5 py-1 bg-blue-50 text-blue-800 font-bold border border-blue-200">Grade A (Excellent)</span>
					</div>
					<div class="flex justify-between items-center text-xs">
						<span class="font-bold text-slate-700">30 – 44 Points</span>
						<span class="px-2.5 py-1 bg-amber-50 text-amber-800 font-bold border border-amber-200">Grade B (Very Good)</span>
					</div>
					<div class="flex justify-between items-center text-xs">
						<span class="font-bold text-slate-700">15 – 29 Points</span>
						<span class="px-2.5 py-1 bg-purple-50 text-purple-800 font-bold border border-purple-200">Grade C (Good)</span>
					</div>
					<div class="flex justify-between items-center text-xs">
						<span class="font-bold text-slate-700">Under 15 Points</span>
						<span class="px-2.5 py-1 bg-slate-50 text-slate-800 font-bold border border-slate-200">Grade D (Satisfactory)</span>
					</div>
				</div>
			</div>

		</section>

		<!-- Divider line -->
		<hr class="border-[#cbd5e1]" />

		<!-- Core Sections Grid -->
		<section class="space-y-12" id="tracks">
			<div class="max-w-2xl space-y-3">
				<h2 class="text-2xl sm:text-3xl font-bold text-[#0f172a] text-serif">
					Cell Foundation Pillars
				</h2>
				<p class="text-xs text-slate-500 font-medium font-sans">
					Through peer learning, campus representation, and social responsibility, we guide students to construct robust portfolios.
				</p>
			</div>

			<!-- Three columns editorial layout -->
			<div class="grid grid-cols-1 md:grid-cols-3 gap-8">
				
				<!-- Column 1 -->
				<div class="space-y-4 border-l border-slate-300 pl-6 py-2">
					<span class="text-2xl font-normal text-slate-400 text-serif">01</span>
					<h3 class="text-lg font-bold text-slate-900 text-serif">Structured Skill Tracks</h3>
					<p class="text-xs text-slate-600 leading-relaxed font-sans">
						Participate in verified self-learning paths. We cover Technical Web Systems, Public Speaking & Oratory, NSS Social Service, and Athletics.
					</p>
				</div>

				<!-- Column 2 -->
				<div class="space-y-4 border-l border-slate-300 pl-6 py-2" id="activities">
					<span class="text-2xl font-normal text-slate-400 text-serif">02</span>
					<h3 class="text-lg font-bold text-slate-900 text-serif">Scheduled Cell Activities</h3>
					<p class="text-xs text-slate-600 leading-relaxed font-sans">
						Register for campus activities, including hackathons, debating forums, blood donation camps, and inter-department sports meets.
					</p>
				</div>

				<!-- Column 3 -->
				<div class="space-y-4 border-l border-slate-300 pl-6 py-2" id="leaderboard">
					<span class="text-2xl font-normal text-slate-400 text-serif">03</span>
					<h3 class="text-lg font-bold text-slate-900 text-serif">Student Score Registry</h3>
					<p class="text-xs text-slate-600 leading-relaxed font-sans">
						Submit participation certificates and verified logs directly to the cell board to update your extra-curricular tally and view academic leaderboards.
					</p>
				</div>

			</div>
		</section>

		<!-- Divider line -->
		<hr class="border-[#cbd5e1]" />

		<!-- Verification/Grading Detailed section -->
		<section class="grid grid-cols-1 lg:grid-cols-12 gap-12 py-6" id="grading">
			<div class="lg:col-span-5 space-y-6">
				<h2 class="text-2xl sm:text-3xl font-bold text-[#0f172a] text-serif">
					Accreditation & <br />
					Grading Scheme
				</h2>
				<p class="text-xs text-slate-600 leading-relaxed font-sans">
					All student certificates and service logs are audited by the departmental committee before final grades are submitted to the university.
				</p>
				<div class="p-5 border-l-4 border-[#991b1b] bg-white border border-[#cbd5e1] rounded-none">
					<p class="text-xs text-slate-600 italic font-sans font-medium">
						"Personality development is not merely an addition; it is an essential component of professional readiness."
					</p>
					<p class="text-[10px] font-bold text-slate-500 mt-2 uppercase tracking-wide">
						— DAVV Board of Studies Guidelines
					</p>
				</div>
			</div>

			<div class="lg:col-span-7 space-y-6">
				<div class="bg-white border border-[#cbd5e1] p-6 space-y-4">
					<h3 class="text-base font-bold text-slate-900 text-serif border-b border-[#e2e8f0] pb-2">
						Point Allocation Rules
					</h3>
					<div class="space-y-3 font-sans">
						<div class="flex items-start gap-3">
							<span class="text-[#991b1b] font-bold mt-0.5">▪</span>
							<p class="text-xs text-slate-600 leading-relaxed">
								<strong>National / State Representation:</strong> Up to 20 points per event for representing DAVV in national debate matches or athletic tournaments.
							</p>
						</div>
						<div class="flex items-start gap-3">
							<span class="text-[#991b1b] font-bold mt-0.5">▪</span>
							<p class="text-xs text-slate-600 leading-relaxed">
								<strong>College Hackathons & Project Meets:</strong> 5 to 15 points depending on rank and tier of involvement.
							</p>
						</div>
						<div class="flex items-start gap-3">
							<span class="text-[#991b1b] font-bold mt-0.5">▪</span>
							<p class="text-xs text-slate-600 leading-relaxed">
								<strong>Social Service & Volunteering (NSS/NCC):</strong> 1 point per 10 hours of verified field work, up to a maximum of 25 points per semester.
							</p>
						</div>
					</div>
				</div>
			</div>
		</section>

		<!-- Portal Entrance Callout -->
		<section class="bg-slate-900 text-white p-12 text-center space-y-6 rounded-none relative overflow-hidden">
			<div class="absolute inset-0 bg-[#0f172a] opacity-95"></div>
			<div class="relative z-10 max-w-xl mx-auto space-y-4">
				<h3 class="text-2xl font-bold text-serif">Student Registry Portal</h3>
				<p class="text-xs text-slate-400 font-sans leading-relaxed">
					Access your student profile, upload credentials, view points status, and verify your semester extra-curricular grades.
				</p>
				<div class="pt-4">
					<button
						onclick={() => alert("Portal entrance triggers login prompt.")}
						class="px-8 py-3 bg-[#991b1b] hover:bg-red-800 text-white text-xs font-bold uppercase tracking-wider rounded-none inline-block transition-colors border-0 cursor-pointer font-sans"
					>
						Enter Portal
					</button>
				</div>
			</div>
		</section>

	</main>

	<!-- Footer -->
	<footer class="bg-white border-t border-[#cbd5e1] py-8 text-center text-xs text-slate-500 font-sans">
		<div class="max-w-7xl mx-auto px-6">
			<p>© 2026 Institute of Infrastructure & Professional Studies (IIPS), DAVV. All rights reserved.</p>
			<p class="mt-1 text-[10px] text-slate-400">ISPARC — Skill, Personality Advancement & Refinement Cell</p>
		</div>
	</footer>
</div>
