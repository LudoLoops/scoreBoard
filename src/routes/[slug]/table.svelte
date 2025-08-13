<script lang="ts">
	const { users } = $props();
	let renderList = $derived(users);

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
	<table class="table table-xs lg:table-md">
		<thead>
			<tr>
				<th onclick={() => sort('userId')} class="hidden cursor-pointer text-left capitalize">Id</th
				>
				<th onclick={() => sort('nom')} class="cursor-pointer text-left capitalize">👽 Nom</th>
				<th onclick={() => sort('note')} class="cursor-pointer text-center capitalize">💫 note</th>
				<th onclick={() => sort('bonus')} class="cursor-pointer text-center capitalize">💥 bonus</th
				>
				<th onclick={() => sort('xp')} class="cursor-pointer text-center capitalize">🚀 xp</th>
				<th class="hidden text-center capitalize sm:block">tags</th>
			</tr>
		</thead>
		<tbody>
			{#each renderList as staff}
				<tr class="hover:bg-base-300/20">
					<th class="hidden">
						{staff.userId}
					</th>
					<td class="font-bold">
						{staff.nom}
					</td>
					<td>
						<div class="flex flex-col items-center justify-center gap-4">
							<span class="text-primary">
								{staff.note || 0}
							</span>
							<progress
								class="progress md:w-56"
								class:progress-success={staff.note >= 40}
								class:progress-secondary={staff.note >= 30}
								class:progress-warning={staff.note >= 26 && staff.note <= 30}
								class:progress-error={staff.note >= 0 && staff.note < 26}
								value={staff.note || 0}
								max="50"
							></progress>
						</div>
					</td>
					<td class="text-center">
						<div class="badge badge-outline badge-primary">{staff.bonus || 0}</div>
					</td>
					<td>
						<div class="flex flex-col items-center justify-center">
							<div
								class="radial-progress"
								class:progress-success={staff.xp >= 50}
								class:progress-secondary={staff.xp >= 40}
								class:progress-warning={staff.xp >= 30 && staff.xp <= 40}
								class:progress-error={staff.xp >= 0 && staff.xp < 30}
								style="--value:{staff.xp || 0}; --size:4rem;"
								aria-valuenow={staff.xp || 0}
								role="progressbar"
							>
								<span class="text-bold text-primary">
									{staff.xp || 0}
								</span>
							</div>
						</div>
					</td>
					<td>
						<div class="hidden flex-wrap gap-2 sm:flex">
							{#each staff.tags as tag}
								<span class="badge badge-soft">
									{tag}
								</span>
							{/each}
						</div>
					</td>
				</tr>
			{/each}
		</tbody>
	</table>
</div>
