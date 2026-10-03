<script lang="ts">
	import type { Snippet } from "svelte";

	interface Props {
		header: string;
		subtitle: string;
		children: Snippet;
	}

	let { header, subtitle, children }: Props = $props();
</script>

<div
	class="flex flex-wrap text-center sm:flex-row p-2 sm:pl-8 sm:justify-between md:text-left md:gap-20 md:mt-6 items-center"
>
	<div class="flex flex-col max-sm:mt-4 max-w-150 gap-2">
		<h1 class="text-xl">{header}</h1>
		<p>
			{@html subtitle}
		</p>
	</div>

	{@render children()}
</div>
