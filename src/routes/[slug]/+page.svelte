<script lang="ts">
	import Table from './table.svelte';

	const { data } = $props();

	const { projet } = data;

	let recherche = $state('');
	let userList = projet.users;

	let users = $state(userList);

	$effect(() => {
		users = search(recherche);
	});

	function search(terme: string) {
		if (!terme) return userList; // Si vide, tout afficher

		const termeNet = terme.toLowerCase().trim();

		return userList.filter((item) => {
			return Object.values(item).some((valeur) => {
				if (Array.isArray(valeur)) {
					// Si c'est un tableau (ex: tags), on cherche dedans
					return valeur.some((tag) => tag.toString().toLowerCase().includes(termeNet));
				} else {
					// Sinon, on convertit en string et on cherche
					return valeur.toString().toLowerCase().includes(termeNet);
				}
			});
		});
	}
</script>

<header class="navbar flex justify-between bg-base-100/20 shadow-sm">
	<a href="/" class="gradient-btn btn text-xl hover:bg-neutral">Dashboard XP Pirate™</a>

	<label class="input input-primary">
		<svg class="h-[1em] opacity-50" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
			<g
				stroke-linejoin="round"
				stroke-linecap="round"
				stroke-width="2.5"
				fill="none"
				stroke="currentColor"
			>
				<circle cx="11" cy="11" r="8"></circle>
				<path d="m21 21-4.3-4.3"></path>
			</g>
		</svg>
		<input type="search" bind:value={recherche} required placeholder="Rechercher" class="" />
	</label>
</header>

<h2 class=" text-neon my-8 text-center text-2xl">
	🏴‍☠️ {projet.name}
</h2>

<Table {users} />

