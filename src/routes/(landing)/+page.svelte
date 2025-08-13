<script lang="ts">
	import Head from '$components/Head.svelte';

	import equipage from '$content/equipage.json';
	import Card from '$ui/card.svelte';
</script>

<Head />

<h1 class="text-center text-4xl">Dashboard XP Pirate™</h1>

<div>
	<h2 class="text-2xl capitalize underline">Projets :</h2>
	<ul class="flex flex-col gap-2">
		{#each equipage as projet}
			<li>
				<a href="/{projet.slug}"> {projet.name}</a>
			</li>
		{/each}
	</ul>
</div>

<Card></Card>
