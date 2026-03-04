<script>
  let { label, required, value = $bindable() } = $props();
</script>

<div class="flex items-center gap-3 pt-2">
  <label class="text-sm font-medium text-gray-700 whitespace-nowrap">
    {label}
    {#if required}
      <span class="text-red-500">*</span>
    {/if}
  </label>

  <input
    type="text"
    bind:value
    {required}
    class="w-full rounded-md border border-gray-300 px-3 py-2 text-sm shadow-sm focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500"
  />
</div>
