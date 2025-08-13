<script lang="ts">
	import { onMount } from 'svelte';

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
				<th onclick={() => sort('nom')} class="cursor-pointer text-left capitalize">Nom</th>
				<th onclick={() => sort('note')} class="text-centercursor-pointer text-center capitalize"
					>note</th
				>
				<th onclick={() => sort('bonus')} class="cursor-pointer text-center capitalize">bonus</th>
				<th onclick={() => sort('xp')} class="cursor-pointer text-center capitalize">xp</th>
				<th class="text-center capitalize">tags</th>
			</tr>
		</thead>
		<tbody>
			{#each renderList as staff}
				<tr class="hover:bg-base-300/20">
					<th>{staff.userId}</th>
					<td>{staff.nom}</td>
					<td>
						<div class="flex flex-col items-center justify-center gap-4">
							<span class="text-primary">
								{staff.note}
							</span>
							<progress
								class="progress w-56 progress-secondary"
								class:progress-success={staff.note >= 40}
								class:progress-neutral={staff.note >= 30}
								class:progress-warning={staff.note >= 26 && staff.note <= 30}
								class:progress-error={staff.note >= 0 && staff.note < 26}
								value={staff.note}
								max="50"
							></progress>
						</div>
					</td>
					<td class="text-center">
						<div class="badge badge-outline badge-primary">{staff.bonus}</div>
					</td>
					<td>
						<div class="flex flex-col items-center justify-center">
							<div
								class="radial-progress"
								class:progress-success={staff.xp >= 50}
								class:progress-neutral={staff.xp >= 40}
								class:progress-warning={staff.xp >= 30 && staff.xp <= 40}
								class:progress-error={staff.xp >= 0 && staff.xp < 30}
								style="--value:{staff.xp};"
								aria-valuenow={staff.xp}
								role="progressbar"
							>
								<span class="text-bold text-primary">
									{staff.xp}
								</span>
							</div>
						</div>
					</td>
					<td>{staff.tags}</td>
				</tr>
			{/each}
		</tbody>
	</table>
</div>
