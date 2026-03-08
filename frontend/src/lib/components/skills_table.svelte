<script>
  let { skills } = $props();

  let categories = $derived.by(() => {
    return [...new Set(skills.map(skill => skill.category))];
  });

  let selectedCategory = $derived.by(() => {
    if (categories.length > 0) {
      return categories[0]
    }
  })

  let filteredSkills = $derived.by(() => {
    return skills.filter(skill => skill.category === selectedCategory);
  });
</script>

<section class="py-8 w-full overflow-x-auto">
  <table class="border-collapse border border-black w-full max-w-7xl lg:max-w-xl bg-skills mx-auto">
    <thead class="bg-gray-100 text-center">
      <tr>
        <th class="border border-black px-4 py-2 text-center">
          <p>Skills:</p>
        </th>
        <th class="border border-black px-4 py-2 text-center">
          <label for="categories">Category:</label>
          <select
            name="categories"
            id="categories"
            class="text-center"
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
          <td class="border border-black px-4 py-2 text-center" colspan="2">{skill.name}</td>
        </tr>
      {/each}
    </tbody>
  </table>
</section>

