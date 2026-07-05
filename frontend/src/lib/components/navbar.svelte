<script>
	import { routes } from "$lib/routes/routes.svelte";
</script>

{#snippet listNavItem(name, fn)}
	<li class="flex w-full h-full">
		<a
			class="flex items-center text-sm cursor-pointer w-full h-full px-12 border-r hover:bg-secondary hover:text-font transition"
			href={fn()}
		>
			{name}
		</a>
	</li>
{/snippet}

<nav class="flex justify-between bg-tertiary text-font border-b h-16">
	<ul class="flex h-full items-center">
		{#each Object.entries(routes) as [name, fn]}
			{@render listNavItem(name, fn)}
		{/each}
	</ul>
</nav>
