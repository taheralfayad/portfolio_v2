<script>
	let { label, required, value = $bindable() } = $props();
</script>

<div>
	<label class="block">
		{label}
		{#if required}<span class="text-red-500">*</span>{/if}
	</label>
	<textarea
		rows="4"
		{required}
		bind:value
		class="mt-1 w-full px-3 py-2 bg-tertiary"
	></textarea>
</div>
