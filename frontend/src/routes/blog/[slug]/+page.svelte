<script lang="ts">
	import { onMount } from "svelte";
	import { page } from "$app/state";

	import Pill from "$lib/components/home/pill.svelte";
	import Stars from "$lib/components/home/blog/stars.svelte";
	import Markdown from "$lib/design-system/markdown.svelte";
	import LoadingSpinner from "$lib/components/home/blog/loading_spinner.svelte";

	import { api } from "$lib/utils/api.svelte.js";
	import { normalizeDate } from "$lib/utils/utils.svelte";

	import type { Blog, BlogResponse } from "../types.ts";

	let blog: Blog | null = $state(null);
	let isLoading = $state(true);

	onMount(async () => {
		try {
			const resp: BlogResponse = await api.get(
				`/blogs/${page.params.slug}`,
			);

			blog = {
				content: resp.content,
				date: normalizeDate(resp.created_at),
				metadata: JSON.parse(resp.metadata),
			};
		} catch (err) {
			console.error(err);
			blog = null;
		} finally {
			isLoading = false;
		}
	});

	$inspect(blog);
</script>

<section class="flex flex-col gap-6 p-5 min-h-screen w-full">
	{#if isLoading}
		<LoadingSpinner />
	{:else}
		<a href="/blog" class="text-sm underline w-fit hover:cursor-pointer">
			&larr; Back to all posts
		</a>

		{#if blog}
			<article class="flex flex-col gap-4 max-w-3xl">
				<div class="flex flex-col gap-2">
					<div class="flex flex-col">
						<div class="flex flex-row items-center gap-4">
							<h1 class="text-3xl font-bold">
								{blog.metadata.title}
							</h1>
							{#if blog.metadata.creator}
								<p class="text-sm italic">
									{blog.metadata.creator}
								</p>
							{/if}
							{#if blog.metadata.rating}
								<Stars rating={blog.metadata.rating} />
							{/if}
						</div>

						{#if blog.date}
							<p class="text-sm mt-2">{blog.date}</p>
						{/if}
					</div>

					{#if blog.metadata.clickbait}
						<p class="text-md">{blog.metadata.clickbait}</p>
					{/if}

					{#if blog.metadata.tags}
						<div class="flex gap-2">
							{#each blog.metadata.tags as tag}
								<div>
									<Pill label={tag} dismissible={false} />
								</div>
							{/each}
						</div>
					{/if}
				</div>

				<hr class="w-[50%]" />

				<Markdown content={blog.content} class="mt-2" />
			</article>
		{:else}
			<p>Post not found.</p>
		{/if}
	{/if}
</section>
