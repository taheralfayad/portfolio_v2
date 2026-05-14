<script>
  let { label, required, value = $bindable(), onFocus } = $props();
</script>

<div class="flex items-center gap-3 pt-2">
  <label class="text-sm font-medium whitespace-nowrap" for={label}>
    {label}
    {#if required}
      <span class="text-red-500">*</span>
    {/if}
  </label>

  <input
    id={label}
    type="text"
    bind:value
    {required}
    onfocus={onFocus}
    class="w-full border border-black px-3 py-2 text-sm shadow-sm focus:border-experience focus:outline-none focus:ring-1 focus:ring-experience"
  />
</div>
