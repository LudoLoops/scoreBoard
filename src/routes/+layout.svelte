<script lang="ts">
	import '../app.css';
	import Navbar from '$c/Navbar.svelte';
	import Breakpoint from '$debug/breakpoint.svelte';

	let { children } = $props();
	let contentEl: HTMLElement | null = null;

	const updateNavbarHeight = (height: number) => {
		if (contentEl) {
			contentEl.style.setProperty('--navbar-height', `${height}px`);
		}
	};
</script>

<Navbar onResize={updateNavbarHeight} />
<main bind:this={contentEl} class="pt-[var(--navbar-height)]">
	{@render children()}
</main>

<Breakpoint />
