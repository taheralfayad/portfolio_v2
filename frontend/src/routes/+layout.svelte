<script>
	import Navbar from "$lib/components/navbar.svelte";
	import "./layout.css";

	let { children } = $props();
</script>

<main class="bg-background text-font">
	<Navbar />
	{@render children()}
</main>
