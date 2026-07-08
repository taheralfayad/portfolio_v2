<script>
	let { title, metric } = $props();
</script>

<div
	class="relative flex flex-col items-center justify-center p-6 w-full min-w-70 max-w-sm
        bg-secondary"
>
	{#if title}
		<span class="text-lg font-bold">{title}</span>
	{/if}

	{#if metric}
		<span class="text-5xl font-bold mt-2">{metric}</span>
	{/if}
</div>
