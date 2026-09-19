<script>
	import BigInput from "$lib/design-system/big_input.svelte";
	import Markdown from "$lib/design-system/markdown.svelte";

	let metadata = $state(
		`{
			"title": "",
			"author": "",
			"rating": 5,
			"tags": ["", "", ""]
		}`,
	);
	let pageContent = $state("");

	let metadataFormatted = $derived.by(() => {
		try {
			const jsonObject = JSON.parse(metadata);
			return JSON.stringify(jsonObject, null, 2);
		} catch (e) {
			return e;
		}
	});

	$inspect(pageContent);
</script>

<div class="flex flex-col w-full h-full items-start justify-start gap-10 p-10">
	<div class="flex flex-row items-center justify-start w-full gap-10">
		<BigInput label="Metadata" bind:value={metadata} required={true} />
		<pre>{metadataFormatted}</pre>
	</div>
	<div class="flex flex-row items-start justify-center w-full gap-10">
		<textarea
			class="flex-1 min-h-screen bg-neutral-300 text-black p-4"
			placeholder="content go here pal"
			bind:value={pageContent}
		></textarea>
		<Markdown
			content={pageContent}
			class="flex-1 min-h-screen overflow-auto border border-neutral-300 p-4"
		/>
	</div>
</div>
