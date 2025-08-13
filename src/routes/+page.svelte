<script lang="ts">
	import Head from '$components/Head.svelte';

	import equipage from '$content/equipage.json';
</script>

<Head />

<h1>Dashboard XP Pirate™</h1>

<div class="flex">
	<ul>
		{#each equipage as projet}
			<li>
				<a href="/{projet.slug}">Projet: {projet.name}</a>
			</li>
		{/each}
	</ul>
</div>
