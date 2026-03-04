<script>
  let { data } = $props();

  let headers = $derived.by(() => {
    if (Array.isArray(data) && data.length > 0) {
      return Object.keys(data[0]);
    }
    return [];
  });

  let rows = $derived.by(() => {
    if (Array.isArray(data)) {
      return data.map((item) => Object.values(item));
    }
    return Object.entries(data).map(([k, v]) => [
      k,
      typeof v === "object" ? JSON.stringify(v) : v,
    ]);
  });

  let plainObjectHeaders = $derived.by(() => {
    if (!Array.isArray(data)) return ["Key", "Value"];
    return headers;
  });
</script>

<section class="flex flex-wrap justify-center py-8">
  <table
    class="border-collapse border border-black max-w-7xl min-w-[500px] bg-skills"
  >
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
