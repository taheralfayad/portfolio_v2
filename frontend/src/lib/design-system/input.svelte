<script>
  let { label, required, value = $bindable(), onFocus } = $props();
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
    onfocus={onFocus}
    class="w-full border border-black px-3 py-2 text-sm shadow-sm focus:border-experience focus:outline-none focus:ring-1 focus:ring-experience"
  />
</div>
