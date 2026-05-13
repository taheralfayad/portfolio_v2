<script>
  let {
    label,
    required,
    value = $bindable(),
    options = [],
    onchange,
  } = $props();
</script>

<div class="flex items-center gap-3 pt-2">
  <label class="text-sm font-medium whitespace-nowrap">
    {label}
    {#if required}
      <span class="text-red-500">*</span>
    {/if}
  </label>
  <select
    bind:value
    {required}
    {onchange}
    class="w-full min-w-[180px] border border-black px-3 py-2 text-sm shadow-sm focus:border-experience focus:outline-none focus:ring-1 focus:border-experience"
  >
    {#each options as option}
      <option class="border-black" value={option.value}>{option.label}</option>
    {/each}
  </select>
</div>
