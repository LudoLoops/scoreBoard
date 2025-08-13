<script lang="ts">
	import { onMount, tick } from 'svelte';
	import { AppBar } from '@skeletonlabs/skeleton-svelte';
	import ArrowLeft from '@lucide/svelte/icons/arrow-left';
	import Calendar from '@lucide/svelte/icons/calendar';
	import CircleUser from '@lucide/svelte/icons/circle-user';
	import Menu from '@lucide/svelte/icons/menu';

	import LightSwitch from '$c/LightSwitch.svelte';

	export let onResize;

	let el: HTMLElement;

	onMount(() => {
		const update = () => {
			const height = el.offsetHeight;
			onResize?.(height);
		};

		const obs = new ResizeObserver(update);
		obs.observe(el);
		update();

		return () => obs.disconnect();
	});
</script>

<div bind:this={el} class="fixed w-full z-999">
	<AppBar
		headlineClasses=" sm:hidden"
		centerClasses="hidden sm:block"
		background="bg-primary-200-800"
	>
		{#snippet lead()}
			<button onclick={() => history.back()}>
				<ArrowLeft size={24} />
			</button>
		{/snippet}

		{#snippet trail()}
			<div class="hidden items-center space-x-4 sm:flex">
				<LightSwitch />
				<Calendar size={20} />
				<CircleUser size={20} />
			</div>
			<div class="block sm:hidden">
				<Menu size={20} />
			</div>
		{/snippet}

		{#snippet headline()}
			<h2 class="h2">Svelte Forge</h2>
		{/snippet}
		<a href="/">
			<span>Svelte Forge</span>
		</a>
	</AppBar>
</div>
