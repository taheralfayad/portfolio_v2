<script>
	import FilterBar from "$lib/components/filter_bar.svelte";

	let searchText = $state("");
	let selectedOptions = $state([]);

	const tagOptions = [
		{ value: "javascript", label: "JavaScript" },
		{ value: "svelte", label: "Svelte" },
		{ value: "css", label: "CSS" },
		{ value: "devops", label: "DevOps" },
		{ value: "career", label: "Career" },
	];
</script>

<section class="p-5 h-screen">
	<FilterBar
		bind:searchText
		bind:selectedOptions
		options={tagOptions}
		searchPlaceholder="Search posts..."
		pillboxLabel="Tags"
	/>
</section>
