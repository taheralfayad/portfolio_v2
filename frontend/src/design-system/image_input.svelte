<script>
  let { label, onchange } = $props();
</script>

<label class="block text-sm font-medium text-gray-700">
  {label} <span class="text-red-500">*</span>
</label>
<input 
  accept="image/png, image/jpeg"
  type="file"
  {onchange}
  class="mt-1 w-full rounded-md border border-gray-300 px-3 py-2 text-sm shadow-sm focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500"
/>
