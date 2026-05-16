<script>
  import { Table } from "@lucide/svelte";

  let { table, schema } = $props();
</script>

<details>
  <summary class="list-none flex flex-row cursor-pointer pr-6 gap-4">
    <Table />
    {table}
  </summary>
  <ul>
    {#each schema as scheme}
      <li class="my-2">
        <b>{scheme.column_name}</b> - {scheme.column_type}
      </li>
    {/each}
  </ul>
</details>
