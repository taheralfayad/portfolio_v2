<script>
	let { submitForm, title, children, editMode, editHook, data } = $props();
	$inspect(data);
</script>

{#snippet dataPreview(data)}
	<section
		class="min-w-xl max-w-xl max-h-150 overflow-y-auto space-y-6 bg-secondary p-8"
	>
		{#each data as datum}
			{#each Object.entries(datum) as [key, value]}
				<p><strong>{key}:</strong> {value}</p>
			{/each}
			<button
				type="submit"
				class="inline-flex items-center justify-center px-4 py-2 font-semibold transition disabled:cursor-not-allowed disabled:opacity-60 hover:cursor-pointer"
				onclick={() => editHook(datum)}
			>
				Edit this entry
			</button>
			<hr />
		{/each}
	</section>
{/snippet}

<div class="flex-wrap sm:flex h-full w-full gap-2">
	<form
		onsubmit={submitForm}
		class="min-w-xl max-w-xl space-y-6 bg-secondary p-8"
	>
		{#if editMode}
			<div>
				<button class="text-fuchsia-950" onclick={() => editHook(null)}>
					you are now in edit mode. i love you. click this to exit
					edit mode
				</button>
			</div>
		{/if}
		<h2 class="text-xl font-semibold">
			{title}
		</h2>

		{@render children()}
	</form>

	{@render dataPreview(data)}
</div>
