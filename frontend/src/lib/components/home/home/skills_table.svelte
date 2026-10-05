<script lang="ts">
	import { type Skill } from "$lib/types/skill.svelte";

	interface Props {
		skills: Skill[];
	}

	let { skills }: Props = $props();

	let categories = $derived.by(() => {
		return [...new Set(skills.map((skill) => skill.category))];
	});

	let selectedCategory = $derived.by(() => {
		if (categories.length > 0) {
			return categories[0];
		}
	});

	let filteredSkills = $derived.by(() => {
		return skills.filter((skill) => skill.category === selectedCategory);
	});
</script>

<section class="h-80 w-full overflow-x-auto">
	<table class="max-w-xl lg:max-w-xl">
		<thead>
			<tr>
				<th class="flex gap-2 items-center">
					<label for="categories">Filter by Category:</label>
					<select
						name="categories"
						id="categories"
						class="p-2 bg-secondary"
						bind:value={selectedCategory}
					>
						{#each categories as category}
							<option value={category}>{category}</option>
						{/each}
					</select>
				</th>
			</tr>
		</thead>
		<tbody>
			{#each filteredSkills as skill}
				<tr class="bg-secondary">
					<td class="px-4 py-2">{skill.name}</td>
				</tr>
			{/each}
		</tbody>
	</table>
</section>
