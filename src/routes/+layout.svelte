<script lang="ts">
	import '../app.css';
	import Breakpoint from '$debug/breakpoint.svelte';

	let { children } = $props();
</script>

<main>
	{@render children()}
</main>

<Breakpoint />
