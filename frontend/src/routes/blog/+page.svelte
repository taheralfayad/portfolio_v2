<script>
	import FilterBar from "$lib/components/home/filter_bar.svelte";
	import Pill from "$lib/components/home/pill.svelte";

	let searchText = $state("");
	let selectedOptions = $state([]);

	const tagOptions = [
		{ value: "javascript", label: "JavaScript" },
		{ value: "svelte", label: "Svelte" },
		{ value: "css", label: "CSS" },
		{ value: "devops", label: "DevOps" },
		{ value: "career", label: "Career" },
	];

	const posts = [
		{
			metadata: {
				id: "svelte-portfolio",
				title: "Building a Svelte Portfolio from Scratch",
				date: "05/05/01",
				clickbait:
					"Why Svelte is the perfect framework for personal sites.",
				tags: [
					{ value: "svelte", label: "Svelte" },
					{ value: "css", label: "CSS" },
				],
			},
		},
		{
			metadata: {
				id: "js-prototypes",
				title: "JavaScript Prototypes Explained",
				clickbait: "A deep dive into the JS inheritance model.",
				tags: [{ value: "javascript", label: "JavaScript" }],
			},
		},
		{
			metadata: {
				id: "devops-journey",
				title: "My Journey into DevOps",
				clickbait: "From deployment nightmares to CI/CD zen.",
				tags: [
					{ value: "devops", label: "DevOps" },
					{ value: "career", label: "Career" },
				],
			},
		},
		{
			metadata: {
				id: "la-haine-review",
				title: "Review: La Haine",
				creator: "Matthew Kassovitz",
				rating: 3,
				clickbait: "A must-read for every software engineer.",
				tags: [
					{ value: "career", label: "Career" },
					{ value: "javascript", label: "JavaScript" },
				],
			},
		},
	];

	function filteredPosts() {
		return posts.filter((post) => {
			const matchesSearch =
				searchText === "" ||
				post.metadata.title
					.toLowerCase()
					.includes(searchText.toLowerCase()) ||
				post.metadata.clickbait
					.toLowerCase()
					.includes(searchText.toLowerCase());
			const matchesFilters =
				selectedOptions.length === 0 ||
				post.metadata.tags.some((tag) =>
					selectedOptions.includes(tag.value),
				);
			return matchesSearch && matchesFilters;
		});
	}
</script>

{#snippet blogCard(item)}
	{#snippet stars(rating)}
		{#snippet star(i, rating)}
			<div class={i < rating ? "text-amber-400" : "text-gray-300"}>
				<svg viewBox="0 0 24 24" fill="currentColor" class="h-6 w-6">
					<path
						d="M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z"
					/>
				</svg>
			</div>
		{/snippet}
		<div class="flex">
			{#each Array(5) as _, i}
				{@render star(i, rating)}
			{/each}
		</div>
	{/snippet}

	<div class="flex flex-col gap-4 pb-5">
		<div class="flex flex-col">
			<div class="flex flex-row items-center gap-4">
				<button class="hover:cursor-pointer">
					<a href={`/blog/${item.metadata.id}`}>
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
					{@render stars(item.metadata.rating)}
				{/if}
			</div>

			{#if item.metadata.date}
				<p class="text-sm">{item.metadata.date}</p>
			{/if}
		</div>

		<p>{item.metadata.clickbait}</p>

		{#if item.metadata.tags}
			<div class="flex gap-2">
				{#each item.metadata.tags as tag}
					<div class="max-w-20">
						<Pill label={tag.label} dismissible={false} />
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
			options={tagOptions}
			searchPlaceholder="Search posts..."
			pillboxLabel="Tags"
		/>
	</div>
	<div class="flex-col gap-4">
		{#each filteredPosts() as post}
			{@render blogCard(post)}
		{/each}
	</div>
</section>
