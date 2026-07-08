<script>
	import { ChevronDown, ChevronUp } from "@lucide/svelte";
	import { routes } from "$lib/routes/routes.svelte";

	const BUTTON_STYLES = "pr-2 w-10 h-10";

	let openDropdown = $state(null);

	function hasChildren(fn) {
		return typeof fn === "function" && Object.keys(fn).length > 0;
	}
</script>

{#snippet navLink(name, fn)}
	<a
		class="flex items-center cursor-pointer w-full px-12 hover:bg-secondary hover:text-font transition h-16"
		href={fn()}
	>
		{name}
	</a>
{/snippet}

{#snippet navItem(name, fn)}
	<li class="relative flex w-full h-full border-r">
		{#if hasChildren(fn)}
			<div
				class="flex w-full h-full hover:bg-secondary hover:text-font items-center justify-center cursor-pointer"
				onmouseenter={() => (openDropdown = name)}
				onmouseleave={() => (openDropdown = null)}
				role="group"
			>
				{@render navLink(name, fn)}

				{#if openDropdown === name}
					<ul
						class="absolute top-full left-0 flex flex-col bg-tertiary border min-w-full z-10"
					>
						{#each Object.entries(fn) as [childName, childFn]}
							<li class="flex w-full">
								{@render navLink(childName, childFn)}
							</li>
						{/each}
					</ul>
				{/if}
				{#if openDropdown === name}
					<ChevronUp class={BUTTON_STYLES} />
				{:else}
					<ChevronDown class={BUTTON_STYLES} />
				{/if}
			</div>
		{:else}
			{@render navLink(name, fn)}
		{/if}
	</li>
{/snippet}

<nav class="flex justify-between bg-tertiary text-font border-b h-16">
	<ul class="flex h-full items-center">
		{#each Object.entries(routes) as [name, fn]}
			{@render navItem(name, fn)}
		{/each}
	</ul>
</nav>
