<script lang="ts">
	interface SeoProps {
		title?: string;
		description?: string;
	}

	const { seo, children } = $props<{ seo?: SeoProps }>();

	let { title, description } = seo;

	const siteName = 'svelteForge';

	const seoData = {
		title: title ? `${title} | ${siteName}` : 'SvelteForge — shape your boilerplate with 🔥',
		description:
			description ??
			'⚒️ Pick your stack, 🔨 Strike the anvil.   Forge your perfect SvelteKit boilerplate. Powered by lelab.dev.',
		url: 'https://svelteforge.lelab.dev',
		image: '%sveltekit.assets%/og-image.png',
		locale: 'en'
	};
</script>

<svelte:head>
	<title>{seoData.title}</title>
	<meta property="og:title" content={seoData.title} />
	<meta name="twitter:title" content={seoData.title} />

	<meta property="og:description" content={seoData.description} />

	<meta name="description" content={seoData.description} />
	<meta name="twitter:description" content={seoData.description} />

	<meta property="og:image" content={seoData.image} />
	<meta property="og:url" content={seoData.url} />
	<meta property="og:site_name" content={siteName} />
	<meta property="og:locale" content={seoData.locale} />

	<meta name="twitter:card" content="summary_large_image" />
	<meta name="twitter:title" content={seoData.title} />
	<meta name="twitter:description" content={seoData.description} />
	<meta name="twitter:image" content={seoData.image} />

	{#if children}
		{@render children()}
	{/if}
</svelte:head>
