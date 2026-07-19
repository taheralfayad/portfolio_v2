<script>
	import { ChevronUp } from "@lucide/svelte";
	export let items = [];
</script>

{#snippet experienceCard(item)}
	<div
		class="flex flex-col sm:flex-row items-center justify-center gap-2 sm:gap-6"
	>
		<div
			class="flex flex-row items-center justify-center sm:justify-end gap-2 w-auto sm:w-32 sm:shrink-0 whitespace-nowrap text-md"
		>
			{#if item.startDate}
				<span>{item.startDate}</span>
			{/if}
			<span>—</span>
			{#if item.endDate}
				<span>{item.endDate}</span>
			{:else}
				<span>Now</span>
			{/if}
		</div>
		<div
			class="w-full sm:w-lg max-w-lg bg-secondary p-6 text-center flex flex-col items-center justify-start shrink-0"
		>
			<div>
				<div class="text-lg font-bold leading-tight">{item.title}</div>
				{#if item.subtitle}
					<div class="mt-2 text-md">{item.subtitle}</div>
				{/if}
			</div>
			{#if item.description}
				<div class="mt-3 text-sm leading-relaxed">
					{item.description}
				</div>
			{/if}
		</div>
	</div>
{/snippet}

{#snippet connector()}
	<div class="flex flex-col sm:flex-row justify-center gap-2 sm:gap-6 w-full">
		<div class="hidden sm:block w-32 shrink-0"></div>
		<div
			class="w-full sm:w-lg max-w-lg flex justify-center select-none text-2xl shrink-0"
		>
			<ChevronUp />
		</div>
	</div>
{/snippet}

<div
	class="flex flex-col flex-wrap items-center justify-center gap-6 py-12 px-6 max-w-6xl"
>
	{#each items as item, index}
		{@render experienceCard(item)}
		{#if index < items.length - 1}
			{@render connector()}
		{/if}
	{/each}
</div>
