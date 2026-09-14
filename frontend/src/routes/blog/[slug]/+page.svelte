<script>
	import { page } from "$app/state";
	import { marked } from "marked";

	import Pill from "$lib/components/home/pill.svelte";

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
				title: "Review: A Clockwork Orange",
				creator: "Stanley Kubrick",
				rating: 3,
				clickbait: "A must-read for every software engineer.",
				date: "01-01-2001",
				tags: [
					{ value: "film", label: "Film" },
					{ value: "director", label: "Review" },
					{ value: "director", label: "Review" },
				],
			},
			content: "## hello \n expletive! expletive",
		},
	];

	const post = $derived(
		posts.find((post) => post.metadata.id === page.params.slug),
	);

	const content = $derived(marked(post.content, { breaks: true }));
</script>

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

<section class="flex flex-col gap-6 p-5 min-h-screen w-full">
	<a href="/blog" class="text-sm underline w-fit hover:cursor-pointer">
		&larr; Back to all posts
	</a>

	{#if post}
		<article class="flex flex-col gap-4 max-w-3xl">
			<div class="flex flex-col gap-2">
				<div class="flex flex-col">
					<div class="flex flex-row items-center gap-4">
						<h1 class="text-3xl font-bold">
							{post.metadata.title}
						</h1>
						{#if post.metadata.creator}
							<p class="text-sm italic">
								{post.metadata.creator}
							</p>
						{/if}
						{#if post.metadata.rating}
							{@render stars(post.metadata.rating)}
						{/if}
					</div>

					{#if post.metadata.date}
						<p class="text-sm">{post.metadata.date}</p>
					{/if}
				</div>

				{#if post.metadata.clickbait}
					<p class="text-md">{post.metadata.clickbait}</p>
				{/if}

				{#if post.metadata.tags}
					<div class="flex gap-2">
						{#each post.metadata.tags as tag}
							<div class="max-w-20">
								<Pill label={tag.label} dismissible={false} />
							</div>
						{/each}
					</div>
				{/if}
			</div>

			<hr class="w-[50%]" />

			<div
				class="[&_h1]:text-3xl [&_h2]:text-2xl [&_h2]:mb-2 [&_p]:text-indent leading-loose tracking-normal mt-2"
			>
				{@html content}
			</div>
		</article>
	{:else}
		<p>Post not found.</p>
	{/if}
</section>

