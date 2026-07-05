<script>
	import { ChevronRight } from "@lucide/svelte";
	export let items = [];
</script>

<div
	class="flex flex-wrap items-center justify-center gap-6 py-12 px-6 max-w-6xl mx-auto"
>
	{#each items as item, index}
		<div class="relative shrink-0">
			<div
				class="w-64 min-h-70 bg-secondary p-6 text-center flex flex-col items-center justify-center"
			>
				<div class="text-base font-bold leading-tight">
					{item.title}
				</div>
				{#if item.subtitle}
					<div class="mt-2 text-sm">
						{item.subtitle}
					</div>
				{/if}
				{#if item.description}
					<div class="mt-3 text-xs leading-relaxed">
						{item.description}
					</div>
				{/if}
			</div>
		</div>
		{#if index < items.length - 1}
			<div class="hidden sm:block select-none text-2xl shrink-0">
				<ChevronRight />
			</div>
		{/if}
	{/each}
</div>
