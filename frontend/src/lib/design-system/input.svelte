<script>
	let { label, required, value = $bindable(), onFocus } = $props();
</script>

<div class="flex items-center gap-3 pt-2">
	<label class="font-medium whitespace-nowrap" for={label}>
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
		class="w-full px-3 py-2 shadow-sm focus:ring-1 focus:ring-experience bg-tertiary"
	/>
</div>
