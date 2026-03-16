<script>
  let { data } = $props();

  let headers = $derived.by(() => {
    if (Array.isArray(data) && data.length > 0) {
      return Object.keys(data[0]);
    }
    return [];
  });

  let plainObjectHeaders = $derived.by(() => {
    if (!Array.isArray(data)) return ["No", "Cups", "Found"];
    return headers;
  });

  let rows = $derived.by(() => {
    return data ? data.map((item) => Object.values(item)) : [];
  });
</script>

<section class="py-8 w-full px-4 overflow-x-auto">
  <table class="border-collapse border border-black w-full bg-skills">
    <thead class="bg-gray-100 text-center">
      <tr>
        {#each plainObjectHeaders as header}
          <th class="border border-black px-4 py-2 text-center capitalize"
            >{header}</th
          >
        {/each}
      </tr>
    </thead>
    <tbody class="text-center">
      {#each rows as row}
        <tr>
          {#each row as cell}
            <td class="border border-black px-4 py-2">{cell}</td>
          {/each}
        </tr>
      {/each}
    </tbody>
  </table>
</section>
