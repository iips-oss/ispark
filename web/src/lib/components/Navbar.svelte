<script lang="ts">
	import { slide } from 'svelte/transition';

	// Navigation items
	const navItems = [
		{ label: 'Overview', href: '#overview' },
		{ label: 'Skill Tracks', href: '#tracks' },
		{ label: 'Activities', href: '#activities' },
		{ label: 'Leaderboard', href: '#leaderboard' },
		{ label: 'Grading Scheme', href: '#grading' }
	];

	let isOpen = $state(false);

	function toggleMenu() {
		isOpen = !isOpen;
	}
</script>

<header class="w-full bg-white border-b border-[#cbd5e1] sticky top-0 z-50">
	<div class="max-w-7xl mx-auto px-6 sm:px-8">
		<div class="flex items-center justify-between h-20">
			
			<!-- Logo and Brand -->
			<a href="/" class="flex items-center gap-3 group focus:outline-none">
				<div class="flex flex-col">
					<span class="text-2xl font-bold tracking-tight text-slate-900 text-serif">
						i<span class="text-[#991b1b]">SPARC</span>
					</span>
					<span class="text-[9px] font-bold text-slate-500 tracking-widest uppercase mt-0.5 font-sans">
						IIPS DAVV Cell
					</span>
				</div>
			</a>

			<!-- Navigation Links (Desktop) -->
			<nav class="hidden md:flex items-center gap-6" aria-label="Main Navigation">
				{#each navItems as item}
					<a
						href={item.href}
						class="text-xs font-bold text-slate-600 hover:text-[#991b1b] transition-colors font-sans uppercase tracking-wider"
					>
						{item.label}
					</a>
				{/each}
			</nav>

			<!-- Action Buttons (Desktop) -->
			<div class="hidden md:flex items-center gap-4">
				<a
					href="#login"
					class="px-5 py-2.5 text-xs font-bold text-slate-900 border border-slate-300 hover:border-slate-900 transition-colors duration-200 rounded-none font-sans"
				>
					Portal Login
				</a>
			</div>

			<!-- Mobile menu button -->
			<div class="flex md:hidden">
				<button
					onclick={toggleMenu}
					type="button"
					class="text-slate-700 hover:text-slate-900 p-2 focus:outline-none bg-transparent border-0 cursor-pointer"
					aria-label="Toggle Menu"
				>
					<svg class="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
						{#if isOpen}
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
						{:else}
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
						{/if}
					</svg>
				</button>
			</div>

		</div>
	</div>

	<!-- Mobile dropdown menu -->
	{#if isOpen}
		<div transition:slide={{ duration: 150 }} class="md:hidden bg-white border-t border-[#cbd5e1] px-6 py-4 space-y-3">
			<nav class="flex flex-col gap-3">
				{#each navItems as item}
					<a
						href={item.href}
						onclick={() => isOpen = false}
						class="text-sm font-semibold text-slate-600 hover:text-slate-900 py-1"
					>
						{item.label}
					</a>
				{/each}
				<hr class="border-[#e2e8f0] my-2" />
				<a
					href="#login"
					onclick={() => isOpen = false}
					class="w-full text-center px-4 py-2.5 text-sm font-bold text-slate-900 border border-slate-300 hover:border-slate-900 transition-colors rounded-none block"
				>
					Portal Login
				</a>
			</nav>
		</div>
	{/if}
</header>
