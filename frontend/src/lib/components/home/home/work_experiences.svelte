<script lang="ts">
	import { type WorkExperience } from "$lib/types/work_experience.svelte";

	interface Props {
		items: WorkExperience[];
	}

	let { items }: Props = $props();

	$inspect(items);
</script>

{#snippet experience(item: WorkExperience)}
	<div class="flex flex-col gap-3">
		<div>
			<h3 class="text-lg leading-tight">{item.title}</h3>
			{#if item.workplace}
				<h4 class="text-md">{item.workplace}</h4>
			{/if}
			<div>
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
		</div>
		{#if item.description}
			<div class="text-sm leading-relaxed">
				{item.description}
			</div>
		{/if}
	</div>
{/snippet}

<div class="flex flex-col flex-wrap max-w-6xl gap-10">
	{#each items as item}
		{@render experience(item)}
	{/each}
</div>
