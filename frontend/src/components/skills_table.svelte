<script>
  let { skills } = $props();

  let categories = $derived.by(() => {
    return [...new Set(skills.map(skill => skill.category))];
  });

  let selectedCategory = $state(categories[0]);

  let filteredSkills = $derived.by(() => {
    return skills.filter(skill => skill.category === selectedCategory);
  });
</script>

<section class="flex flex-wrap justify-center py-8">
  <table class="border-collapse border border-black max-w-7xl min-w-[500px] bg-skills">
    <thead class="bg-gray-100 text-center">
      <tr>
        <th class="border border-black px-4 py-2 text-center">
          <label for="fruits">Category of skills:</label>
        </th>
        <th class="border border-black px-4 py-2 text-center">
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
          <td class="border border-black px-4 py-2">{skill.name}</td>
        </tr>
      {/each}
      <!-- <tr> -->
      <!--   <td class="border border-black px-4 py-2">Alfreds Futterkiste</td> -->
      <!--   <td class="border border-black px-4 py-2">Alfreds Futterkiste</td> -->
      <!-- </tr> -->
      <!-- <tr> -->
      <!--   <td class="border border-black px-4 py-2">Alfreds Futterkiste</td> -->
      <!-- </tr> -->
    </tbody>
  </table>
</section>

