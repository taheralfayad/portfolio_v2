<script lang="ts">
	import { goto } from "$app/navigation";

	import Input from "$lib/design-system/input.svelte";
	import FormButton from "$lib/design-system/form_button.svelte";

	import { api } from "$lib/utils/api.svelte.js";

	let userName = $state("");
	let password = $state("");

	const logIn = async () => {
		const payload = {
			name: userName,
			password: password,
		};

		try {
			await api.post("/login", payload);
			await goto("/admin");
		} catch (err) {
			console.error(err);
		}
	};
</script>

<form onsubmit={logIn}>
	<div class="min-h-screen flex items-center justify-center">
		<div>
			<Input label={"Name"} bind:value={userName} />
			<Input label={"Password"} bind:value={password} />
			<div class="pt-4"></div>
			<FormButton loading={false} />
		</div>
	</div>
</form>
