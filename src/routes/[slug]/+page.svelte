<script lang="ts">
	import Table from './table.svelte';

	const { data } = $props();

	const { projet } = data;
</script>

<Table {projet} />
