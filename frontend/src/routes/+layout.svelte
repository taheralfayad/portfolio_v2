<script>
	import Navbar from "$lib/components/navbar.svelte";
	import "./layout.css";

	let { children } = $props();
</script>

<Navbar />
<main class="bg-background text-font p-5">
	{@render children()}
</main>
