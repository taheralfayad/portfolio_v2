import { redirect } from '@sveltejs/kit';
import { isLoggedIn } from '$lib/utils/utils.svelte'

export async function load() {
  if (!(await isLoggedIn())) {
    throw redirect(307, '/login');
  }
}
