<script>
	let { label, onchange } = $props();
</script>

<label class="block font-medium">
	{label} <span class="text-red-500">*</span>
</label>
<input
	accept="image/png, image/jpeg"
	type="file"
	{onchange}
	class="mt-1 w-full px-3 py-2 bg-tertiary"
/>
