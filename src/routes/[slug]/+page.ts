import { error } from '@sveltejs/kit';
import equipage from '$content/equipage.json';

export async function load({ params }) {
	const slug = params.slug;

	const projet = equipage.find((a) => a.slug === slug);
	if (projet?.slug === slug) {
		return { projet };
	}

	throw error(404, 'Article non trouvé');
}
