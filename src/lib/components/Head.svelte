<script lang="ts">
	interface Props {
		title?: string;
		description?: string;
		href?: string;
		main?: boolean;
		children?: any;
	}

	const {
		title,
		children,
		description = 'Mettez la main sur le trésor des évaluations ! 🗺️ Accédez à votre dashboard de notes de projets en un clin d’œil, suivez les performances, repérez les bateaux qui coulent et hissez les projets gagnants vers la gloire. Organisation, transparence et contrôle total — comme un vrai capitaine du succès !',
		href = 'https://ludoapex.fr',
		main = false
	}: Props = $props();

	const siteName = 'Dashboard XP Pirate™';

	let fullTitle = $state<string>(
		"ludoapex - Sites web sur-mesure pour les pro de la nature et l'aventure"
	);

	if (!main) {
		fullTitle = title ? `${title} | ${siteName}` : siteName;
	}
</script>

<svelte:head>
	<title>{fullTitle}</title>

	<link rel="canonical" {href} />
	<meta property="og:title" content={fullTitle} />
	<meta name="twitter:title" content={fullTitle} />

	<meta property="og:description" content={description} />
	<meta name="description" content={description} />
	<meta name="twitter:description" content={description} />

	{#if children}
		{@render children()}
	{/if}
</svelte:head>
