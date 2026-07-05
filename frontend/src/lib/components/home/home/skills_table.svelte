<script>
	let { skills } = $props();

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

<section class="py-8 w-full px-4 overflow-x-auto">
	<table class="w-full max-w-7xl lg:max-w-xl bg-secondary mx-auto">
		<thead class="text-center">
			<tr>
				<th
					class="flex px-4 py-2 text-center gap-2 justify-center items-center border-b border-t"
				>
					<label for="categories">Filter by Category:</label>
					<select
						name="categories"
						id="categories"
						class="text-center p-2"
						bind:value={selectedCategory}
					>
						{#each categories as category}
							<option value={category}>{category}</option>
						{/each}
					</select>
				</th>
			</tr>
		</thead>
		<tbody class="text-center">
			{#each filteredSkills as skill}
				<tr>
					<td class="px-4 py-2 text-center" colspan="2"
						>{skill.name}</td
					>
				</tr>
			{/each}
		</tbody>
	</table>
</section>
