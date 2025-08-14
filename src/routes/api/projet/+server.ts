import { json } from '@sveltejs/kit';

import equipage from '$content/equipage.json';

const projets = equipage.map((item) => item.name);

export function GET() {
	return json(projets);
}

export function POST() {}
