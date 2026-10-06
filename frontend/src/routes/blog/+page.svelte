<script lang="ts">
	interface FilterableTag {
		value: string;
		label: string;
	}

	import { onMount } from "svelte";

	import FilterBar from "$lib/components/home/filter_bar.svelte";
	import Pill from "$lib/components/home/pill.svelte";
	import Stars from "$lib/components/home/blog/stars.svelte";
	import LoadingSpinner from "$lib/components/home/blog/loading_spinner.svelte";

	import { api } from "$lib/utils/api.svelte.js";
	import { normalizeDate } from "$lib/utils/utils.svelte";

	import type { Blog, BlogResponse } from "$lib/types/blog";

	let searchText = $state("");
	let selectedOptions: string[] = $state([]);

	let isLoading = $state(true);

	let blogs: Blog[] = $state([]);
	let tags: FilterableTag[] = $derived(
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
		} finally {
			isLoading = false;
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
			<div
				class="flex flex-col gap-2 sm:flex-row sm:items-center sm:gap-4"
			>
				<a href={`/blog/${item.id}`}>
					<h2 class="text-xl hover:underline">
						{item.metadata.title}
					</h2>
				</a>
				{#if item.metadata.creator || item.metadata.rating}
					<div class="flex items-center gap-3 sm:gap-4">
						{#if item.metadata.creator}
							<p class="text-sm italic">
								{item.metadata.creator}
							</p>
						{/if}
						{#if item.metadata.rating}
							<Stars rating={item.metadata.rating} />
						{/if}
					</div>
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

<section class="flex flex-col gap-4 h-screen w-full">
	{#if isLoading}
		<LoadingSpinner />
	{:else}
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
	{/if}
</section>
