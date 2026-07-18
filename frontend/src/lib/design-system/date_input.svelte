<script>
	let { label, value = $bindable() } = $props();
</script>

<div>
	<label class="block">
		{label}
	</label>
	<input type="date" bind:value class="mt-1 w-full px-3 py-2 bg-tertiary" />
</div>
