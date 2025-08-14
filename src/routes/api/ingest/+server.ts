import { json } from '@sveltejs/kit';

import equipage from '$content/equipage.json';

export function GET() {
	return json(equipage);
}

export const POST = async ({ request }) => {
	try {
		const data = await request.json();

		// Validate data if needed
		if (!data || !data.name) {
			return json({ error: 'Name is required' }, { status: 400 });
		}

		return json({ message: 'Data received successfully!', data: data }, { status: 201 }); // 201 Created
	} catch (error) {
		console.error('Error processing POST request:', error);
		return json({ error: 'Internal server error' }, { status: 500 });
	}
};
