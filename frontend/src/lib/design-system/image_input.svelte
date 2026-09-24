<script lang="ts">
	interface Props {
		label: string;
		onchange(e: Event): void;
	}

	let { label, onchange }: Props = $props();
</script>

<label class="block font-medium">
	{label} <span class="text-red-500">*</span>
	<input
		accept="image/png, image/jpeg"
		type="file"
		{onchange}
		class="mt-1 w-full px-3 py-2 bg-tertiary"
	/>
</label>
