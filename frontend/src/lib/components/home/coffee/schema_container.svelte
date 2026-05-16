<script>
  import SqlSchema from "$lib/components/home/coffee/sql_schema.svelte";

  let { schemas } = $props();
</script>

<div
  class="relative w-72 shrink-0 h-144 overflow-y-auto
  overflow-x-hidden
  border border-black
  max-w-md
  bg-white
"
>
  <div class="sticky top-4 p-4">
    <h3 class="text-lg mb-4">Schema</h3>

    <div class="flex flex-col gap-6">
      {#each schemas as schema}
        <SqlSchema table={schema.table} schema={schema.schema} />
      {/each}
    </div>
  </div>
</div>
