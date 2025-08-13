<script lang="ts">
	let { children } = $props();
</script>

<div class="navbar bg-base-100/40 shadow-sm">
	<a href="/" class="  text-xl hover:bg-neutral">Dashboard XP Pirate™</a>
</div>

<main class="bg-neutral">
	{@render children()}
</main>
