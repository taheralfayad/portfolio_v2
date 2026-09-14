<script>
	import Pill from "$lib/components/home/pill.svelte";

	let {
		searchText = $bindable(""),
		selectedOptions = $bindable([]),
		options = [],
		filterLabel = "Filter",
		searchPlaceholder = "Search...",
		pillboxLabel = "Filters",
	} = $props();

	let dropdownValue = $state("");

	function addOption(value) {
		if (value && !selectedOptions.includes(value)) {
			selectedOptions = [...selectedOptions, value];
		}
		dropdownValue = "";
	}

	function removeOption(option) {
		selectedOptions = selectedOptions.filter((o) => o !== option);
	}

	function getLabel(value) {
		const match = options.find((o) => o.value === value);
		return match ? match.label : value;
	}
</script>

<div class="flex flex-row flex-wrap items-center gap-3">
	<input
		class="bg-tertiary px-3 py-2"
		type="text"
		placeholder={searchPlaceholder}
		bind:value={searchText}
	/>
	<select
		class="bg-tertiary px-3 py-2 focus:outline-none"
		bind:value={dropdownValue}
		onchange={(e) => addOption(e.currentTarget.value)}
	>
		<option value="">{filterLabel}</option>
		{#each options as option}
			<option value={option.value}>{option.label}</option>
		{/each}
	</select>
	{#if selectedOptions.length > 0}
		<div class="flex flex-row items-center flex-wrap gap-2">
			<p>{pillboxLabel}:</p>
			{#each selectedOptions as option}
				<Pill
					label={getLabel(option)}
					onRemove={() => removeOption(option)}
				/>
			{/each}
		</div>
	{/if}
</div>
