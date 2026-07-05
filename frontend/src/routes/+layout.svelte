<script>
	import { goto } from "$app/navigation";

	import Navbar from "$lib/components/navbar.svelte";
	import { routes } from "$lib/routes/routes.svelte";
	import "./layout.css";

	let { children } = $props();

	const navItems = [
		{
			value: "Home",
			header: "Hello! My name is Ahmad Altaher Alfayad.",
			subtitle:
				"I'm a husband, older brother, dog dad, Arsenal fan, and software developer.",
			onClick: () => {
				currNavValue = navItems.find((item) => item.value === "Home");
				goto("?tab=Home", { replaceState: true });
			},
		},
		{
			value: "Coffee",
			header: "Hello! Join me on my coffee journey.",
			subtitle:
				"I love drinking coffee. I've recently started roasting my own, and this is where I document it.",
			onClick: () => {
				currNavValue = navItems.find((item) => item.value === "Coffee");
				goto("?tab=Coffee", { replaceState: true });
			},
		},
	];

	let currNavValue = $state(navItems.find((item) => item.value === "Home"));

	const getCurrNavValue = () => currNavValue.value;
</script>

<Navbar items={navItems} {getCurrNavValue} />
<main class="bg-background text-font font-semibold">
	{@render children()}
</main>
