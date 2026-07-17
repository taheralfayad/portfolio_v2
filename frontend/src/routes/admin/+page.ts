export const ssr = false;

import { redirect } from '@sveltejs/kit';
import { isLoggedIn } from '$lib/utils/utils.svelte';

export const load = async () => {
	const isLoggedInBool = await isLoggedIn();

	if (!isLoggedInBool) {
		redirect(303, '/login');
	}
}

