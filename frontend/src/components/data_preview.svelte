<script>
  let { data, editHook } = $props();
  console.log(data)
</script>

<section
	class="min-w-xl max-w-xl max-h-[600px] overflow-y-auto space-y-6 rounded-2xl border border-gray-200 bg-white p-8 shadow-sm"
>
  {#each data as datum}
    {#each Object.entries(datum) as [key, value]}
      <p><strong>{key}:</strong> {value}</p>
    {/each}
    <button
      type="submit"
      class="inline-flex items-center justify-center rounded-md bg-blue-600 px-4 py-2 text-sm font-semibold text-white shadow-sm transition hover:bg-blue-700 disabled:cursor-not-allowed disabled:opacity-60 hover:cursor-pointer"
      onclick={() => editHook(datum)}
    >
      Edit this entry
    </button>
    <hr/>
  {/each}
</section>
