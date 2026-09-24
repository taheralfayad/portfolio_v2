<script lang="ts">
	interface Props {
		label: string;
		required?: boolean;
		value?: string;
		class?: string;
		rows?: number;
	}

	let {
		label,
		required,
		value = $bindable(),
		class: className = "",
		rows = 4,
	}: Props = $props();
</script>

<div class={className}>
	<label class="block">
		{label}
		{#if required}<span class="text-red-500">*</span>{/if}
	</label>
	<textarea
		{rows}
		{required}
		bind:value
		class="mt-1 w-full px-3 py-2 bg-tertiary"
	></textarea>
</div>