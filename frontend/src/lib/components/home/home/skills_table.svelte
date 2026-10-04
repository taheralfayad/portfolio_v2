<script lang="ts">
	import { ChevronRight, ChevronLeft } from "@lucide/svelte";
	import { type Skill } from "$lib/types/skill.svelte";

	interface Props {
		skills: Skill[];
	}

	let { skills }: Props = $props();

	let activeIndex = $state(0);

	let categorizedSkills = $derived.by(() => {
		const grouped = Object.groupBy(skills, (skill) => skill.category);
		return Object.entries(grouped).sort(([a], [b]) => a.localeCompare(b));
	});

	let visibleCategories = $derived.by(() => {
		const start = Math.max(0, activeIndex - 1);
		return categorizedSkills
			.map((entry, index) => ({ entry, index }))
			.slice(start, activeIndex + 2);
	});
</script>

{#snippet Card(
	categoryWithSkills: [string, Skill[] | undefined],
	active: boolean,
)}
	<div
		class={`bg-secondary flex flex-col justify-start shrink-0 items-center p-4
			transition-all duration-300 gap-10 text-center tracking-widest
			${active ? "min-h-96 w-56" : "h-48 w-24 opacity-85 overflow-clip"}`}
	>
		<h3 class="shrink-0">{categoryWithSkills[0]}</h3>
		<hr />
		{#each categoryWithSkills[1] as skill}
			<p>{skill.name}</p>
		{/each}
	</div>
{/snippet}

{#snippet Buttons()}
	<button
		type="button"
		onclick={() => {
			if (activeIndex > 0) {
				activeIndex--;
			}
		}}
		aria-label="Previous skill"
		class="p-3 hover:cursor-pointer bg-button border border-black"
	>
		<ChevronLeft color="black" />
	</button>
	<button
		type="button"
		onclick={() => {
			if (activeIndex < categorizedSkills.length - 1) {
				activeIndex++;
			}
		}}
		aria-label="Next skill"
		class="p-3 hover:cursor-pointer bg-button border border-black"
	>
		<ChevronRight color="black" />
	</button>
{/snippet}

<div class="flex flex-col gap-4 justify-center items-center">
	<div class="flex flex-row gap-4 justify-center items-center h-125 w-125">
		{#each visibleCategories as { entry, index } (entry[0])}
			{@render Card(entry, index === activeIndex)}
		{/each}
	</div>
	<div class="flex flex-row gap-6">
		{@render Buttons()}
	</div>
</div>
