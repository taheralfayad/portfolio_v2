<script lang="ts">
	import { X } from "@lucide/svelte";

	interface Props {
		label: string;
		onRemove?: () => {};
		dismissible: boolean;
	}
	let { label, onRemove, dismissible = true }: Props = $props();
</script>

<div class="inline-flex items-center gap-2 px-4 py-2 bg-tertiary">
	<p class="text-sm">{label}</p>
	{#if dismissible}
		<button class="shrink-0 cursor-pointer" onclick={onRemove}>
			<X size={14} />
		</button>
	{/if}
</div>
