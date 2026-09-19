<script lang="ts">
	import { marked } from "marked";

	interface Props {
		content: string;
		class?: string;
	}

	let { content, class: className = "" }: Props = $props();

	let html = $derived.by(() => {
		try {
			return marked(content, { breaks: true });
		} catch (e) {
			return String(e);
		}
	});
</script>

<div class={`prose max-w-none leading-loose ${className}`}>
	{@html html}
</div>

<style>
	:global(.prose) {
		--tw-prose-body: var(--color-font);
		--tw-prose-headings: var(--color-font);
		--tw-prose-lead: var(--color-font);
		--tw-prose-links: var(--color-font);
		--tw-prose-bold: var(--color-font);
		--tw-prose-counters: var(--color-font);
		--tw-prose-bullets: var(--color-font);
		--tw-prose-quotes: var(--color-font);
		--tw-prose-captions: var(--color-font);
		--tw-prose-kbd: var(--color-font);
		--tw-prose-code: var(--color-font);
		--tw-prose-pre-code: var(--color-font);
		--tw-prose-hr: color-mix(
			in oklab,
			var(--color-font) 25%,
			var(--color-background)
		);
		--tw-prose-quote-borders: color-mix(
			in oklab,
			var(--color-font) 25%,
			var(--color-background)
		);
		--tw-prose-kbd-shadows: color-mix(
			in oklab,
			var(--color-font) 25%,
			var(--color-background)
		);
		--tw-prose-pre-bg: color-mix(
			in oklab,
			var(--color-font) 10%,
			var(--color-background)
		);
		--tw-prose-th-borders: color-mix(
			in oklab,
			var(--color-font) 25%,
			var(--color-background)
		);
		--tw-prose-td-borders: color-mix(
			in oklab,
			var(--color-font) 25%,
			var(--color-background)
		);
	}
</style>
