<script lang="ts">
	import { onMount } from "svelte";

	import FilterBar from "$lib/components/home/filter_bar.svelte";
	import Pill from "$lib/components/home/pill.svelte";
	import Stars from "$lib/components/home/blog/stars.svelte";

	import { api } from "$lib/utils/api.svelte.js";
	import { normalizeDate } from "$lib/utils/utils.svelte";

	import type { Blog, Tag, BlogResponse } from "./types";

	let searchText = $state("");
	let selectedOptions = $state([]);

	let blogs: Blog[] = $state([]);
	let tags: Tag[] = $derived(
		blogs.flatMap((blog) =>
			blog.metadata.tags.map((tag) => ({
				value: tag,
				label: tag,
			})),
		),
	);

	onMount(async () => {
		try {
			const resp: BlogResponse[] = await api.get("/blogs");

			blogs = resp.map((blog): Blog => {
				return {
					id: blog.id,
					content: blog.content,
					date: normalizeDate(blog.created_at),
					metadata: JSON.parse(blog.metadata),
				};
			});
		} catch (err) {
			console.error(err);
		}
	});

	function filteredBlogs() {
		return blogs.filter((blog) => {
			const matchesSearch =
				searchText === "" ||
				blog.metadata.title
					.toLowerCase()
					.includes(searchText.toLowerCase()) ||
				blog.metadata.clickbait
					.toLowerCase()
					.includes(searchText.toLowerCase());
			const matchesFilters =
				selectedOptions.length === 0 ||
				blog.metadata.tags.some((tag) => selectedOptions.includes(tag));
			return matchesSearch && matchesFilters;
		});
	}
</script>

{#snippet blogCard(item: Blog)}
	<div class="flex flex-col gap-4 pb-5">
		<div class="flex flex-col">
			<div class="flex flex-row items-center gap-4">
				<button class="hover:cursor-pointer">
					<a href={`/blog/${item.id}`}>
						<h2 class="text-xl hover:underline">
							{item.metadata.title}
						</h2>
					</a>
				</button>
				<!-- the following will only be useful in case the article is a review on a piece of media created by someone else !-->
				{#if item.metadata.creator}
					<p class="text-sm italic">{item.metadata.creator}</p>
				{/if}
				{#if item.metadata.rating}
					<Stars rating={item.metadata.rating} />
				{/if}
			</div>

			{#if item.date}
				<p class="text-sm">{item.date}</p>
			{/if}
		</div>

		<p>{item.metadata.clickbait}</p>

		{#if item.metadata.tags}
			<div class="flex gap-2">
				{#each item.metadata.tags as tag}
					<div>
						<Pill label={tag} dismissible={false} />
					</div>
				{/each}
			</div>
		{/if}
		<hr class="w-[50%]" />
	</div>
{/snippet}

<section class="flex flex-col gap-4 p-5 h-screen w-full">
	<div>
		<FilterBar
			bind:searchText
			bind:selectedOptions
			options={tags}
			searchPlaceholder="Search blogs..."
			pillboxLabel="Tags"
		/>
	</div>
	<div class="flex-col gap-4">
		{#each filteredBlogs() as blog}
			{@render blogCard(blog)}
		{/each}
	</div>
</section>
