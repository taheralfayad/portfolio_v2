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
	<label class="font-medium whitespace-nowrap">
		{label}
		{#if required}
			<span class="text-red-500">*</span>
		{/if}
	</label>
	<select
		bind:value
		{required}
		{onchange}
		class="w-full min-w-45 px-3 py-2 focus:outline-none focus:ring-1 bg-tertiary"
	>
		{#each options as option}
			<option value={option.value}>{option.label}</option>
		{/each}
	</select>
</div>
