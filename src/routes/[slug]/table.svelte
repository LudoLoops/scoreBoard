<script lang="ts">
	const { projet } = $props();

	const { users } = projet;

	let renderList = $state(users);

	let toggle = $state({ note: false, xp: false, bonus: false, nom: false });

	function sortDsc(field: string) {
		if (field === 'nom') {
			return (renderList = users.sort((a: { nom: any }, b: { nom: string }) =>
				b.nom.localeCompare(a.nom)
			));
		}
		return (renderList = users.sort(
			(a: { [x: string]: number }, b: { [x: string]: number }) => b[field] - a[field]
		));
	}

	function sortAsc(field: string) {
		if (field === 'nom') {
			return (renderList = users.sort((a: { nom: string }, b: { nom: any }) =>
				a.nom.localeCompare(b.nom)
			));
		}
		return (renderList = users.sort(
			(a: { [x: string]: number }, b: { [x: string]: number }) => a[field] - b[field]
		));
	}

	function sort(field: string) {
		if (toggle[field as keyof typeof toggle]) {
			sortDsc(field);
		} else {
			sortAsc(field);
		}

		toggle[field as keyof typeof toggle] = !toggle[field as keyof typeof toggle];
	}
</script>

<div class="overflow-x-auto">
	<table class="table">
		<thead>
			<tr>
				<th>Id</th>
				<th onclick={() => sort('nom')}> Nom</th>
				<th onclick={() => sort('note')}>note</th>
				<th onclick={() => sort('bonus')}>bonus</th>
				<th onclick={() => sort('xp')}>xp</th>
				<th>tags</th>
			</tr>
		</thead>
		<tbody>
			{#each renderList as staff}
				<tr>
					<th>{staff.userId}</th>
					<td>{staff.nom}</td>
					<td>{staff.note}</td>
					<td>{staff.bonus}</td>
					<td>{staff.xp}</td>
					<td>{staff.tags}</td>
				</tr>
			{/each}
		</tbody>
	</table>
</div>
