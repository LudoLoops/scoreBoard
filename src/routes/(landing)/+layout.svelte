<script lang="ts">
	let { children } = $props();
</script>

<div class="navbar bg-base-100/20 shadow-sm">
	<a href="/" class="gradient-btn btn text-xl">Dashboard XP Pirate™</a>
</div>

{@render children()}
