<script lang="ts">
	import Head from '$components/Head.svelte';

	import equipage from '$content/equipage.json';
</script>

<Head />

<h1 class="text-neon my-8 text-center text-4xl font-black text-primary">Dashboard XP Pirate™</h1>

<ul class="card mx-8 flex flex-col gap-3">
	<h2 class="card-title pt-8 pl-8 text-2xl capitalize">⚡ Projets :</h2>
	<div class="card-body gap-8">
		{#each equipage as projet, i}
			<a
				href="/{projet.slug}"
				class="card-border hover:text-neon neon-border card mx-auto card-body w-xl border-1 border-primary shadow-primary hover:bg-teal-500/10"
			>
				<li class="grid grid-cols-[70px_1fr] border-secondary text-xl font-bold uppercase">
					<span class=" mr-2 badge badge-neutral">
						🏴‍☠️ {i + 1}
					</span>
					{projet.name}
				</li>
			</a>
		{/each}
	</div>
</ul>
<div class="h-8"></div>
