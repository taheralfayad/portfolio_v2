<script>
	let { items, getCurrNavValue } = $props();

	import NavItem from "$lib/design-system/nav_item.svelte";
</script>

<nav class="flex justify-between bg-tertiary text-font border-b h-16">
	<ul class="flex h-full items-center">
		{#each items as item}
			<NavItem
				value={item.value}
				onClick={item.onClick}
				{getCurrNavValue}
			/>
		{/each}
	</ul>
</nav>
