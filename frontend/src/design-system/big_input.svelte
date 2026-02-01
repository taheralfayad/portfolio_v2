<script>
  let {label, value = $bindable()} = $props();
</script>


<div>
  <label class="block text-sm font-medium text-gray-700">
    {label}
  </label>
  <textarea
    rows="4"
    bind:value={value}
    class="mt-1 w-full rounded-md border border-gray-300 px-3 py-2 text-sm shadow-sm focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500"
  />
</div>
