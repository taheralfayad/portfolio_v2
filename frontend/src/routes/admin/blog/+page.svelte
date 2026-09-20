<script lang="ts">
	import { X } from "@lucide/svelte";
	import { onMount } from "svelte";
	import BigInput from "$lib/design-system/big_input.svelte";
	import ImageInput from "$lib/design-system/image_input.svelte";
	import Input from "$lib/design-system/input.svelte";
	import Markdown from "$lib/design-system/markdown.svelte";

	import { uploadImage } from "$lib/utils/utils.svelte";

	const DEFAULT_METADATA = `{
	  "title": "",
	  "author": "",
	  "rating": 5,
	  "tags": ["", "", ""]
	}`;

	let metadata = $state(DEFAULT_METADATA);
	let pageContent = $state("");
	let loaded = $state(false);
	let imageModalOpen = $state(false);
	let imageData = $state("");
	let imageAltText = $state("");

	onMount(() => {
		metadata = localStorage.getItem("blog:metadata") ?? DEFAULT_METADATA;
		pageContent = localStorage.getItem("blog:content") ?? "";
		loaded = true;
	});

	let metadataFormatted = $derived.by(() => {
		try {
			return JSON.stringify(JSON.parse(metadata), null, 2);
		} catch (e) {
			return `Invalid JSON: ${e.message}`;
		}
	});

	const exitImageModal = () => {
		imageModalOpen = false;
		imageData = "";
		imageAltText = "";
	};

	const upload = async () => {
		const resp = await uploadImage(imageAltText, imageData, "blog");

		if (resp.success) {
			exitImageModal();

			pageContent += "\n";

			pageContent += `![${resp.resp.title}](${resp.resp.image})\n`;
		} else {
			console.error(
				"We encountered an error with your upload: ",
				resp.error,
			);
		}

		return;
	};

	$effect(() => {
		if (!loaded) return;
		const value = metadata;
		const timeout = setTimeout(() => {
			try {
				localStorage.setItem("blog:metadata", value);
			} catch (e) {
				console.error("Failed to save metadata", e);
			}
		}, 300);
		return () => clearTimeout(timeout);
	});

	$effect(() => {
		if (!loaded) return;
		const value = pageContent;
		const timeout = setTimeout(() => {
			try {
				localStorage.setItem("blog:content", value);
			} catch (e) {
				console.error("Failed to save content", e);
			}
		}, 300);
		return () => clearTimeout(timeout);
	});
</script>

<svelte:window
	onkeydown={(e) => {
		if (e.key === "Escape" && imageModalOpen) {
			imageModalOpen = false;
		}
	}}
/>

{#if imageModalOpen}
	<div
		onclick={exitImageModal}
		role="presentation"
		class="absolute m-auto p-25 bg-black/25 w-full h-full"
	>
		<button class="cursor-pointer bg-secondary" onclick={exitImageModal}>
			<X />
		</button>
		<div
			role="presentation"
			aria-label="Add image"
			onclick={(e) => e.stopPropagation()}
			class="flex gap-4 bg-secondary"
		>
			<ImageInput
				label="Add image"
				onchange={(e) => {
					const file = e.target.files?.[0];
					if (!file) return;

					const reader = new FileReader();
					reader.onload = () => {
						imageData = reader.result;
					};
					reader.readAsDataURL(file);
				}}
			/>

			<Input
				label="Add alt-text"
				required={true}
				bind:value={imageAltText}
			/>

			<button class="bg-tertiary cursor-pointer" onclick={upload}>
				Add image
			</button>
		</div>

		<img src={imageData} alt={imageAltText} />
	</div>
{/if}

<div class="flex flex-col w-full h-full items-start justify-start gap-10 p-10">
	<div class="flex flex-row items-center justify-start w-full gap-10">
		<BigInput label="Metadata" bind:value={metadata} required={true} />
		<pre>{metadataFormatted}</pre>
	</div>
	<div class="flex flex-row items-start justify-center w-full gap-10">
		<div class="flex-1 min-h-screen">
			<ul class="flex flex-row w-full bg-tertiary p-3">
				<button
					class="cursor-pointer"
					onclick={() => (imageModalOpen = true)}>Add Image</button
				>
			</ul>
			<textarea
				class="w-full min-h-screen bg-neutral-300 text-black p-4"
				placeholder="content go here pal"
				bind:value={pageContent}
			></textarea>
		</div>
		<Markdown
			content={pageContent}
			class="flex-1 min-h-screen overflow-auto border border-neutral-300 p-4"
		/>
	</div>
</div>
