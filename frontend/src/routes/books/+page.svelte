<script>
	import { onMount } from "svelte";

	import { X } from "@lucide/svelte";

	import { api } from "$lib/utils/api.svelte";

	import Hero from "$lib/components/home/hero.svelte";
	import Carousel from "$lib/design-system/carousel.svelte";

	import Content from "$lib/content/home.json";

	let books = $state([]);
	let images = $state([]);
	let searchText = $state("");
	let statusFilters = $state([]);

	let booksFiltered = $derived.by(() => {
		const query = searchText.trim().toLowerCase();
		const statusPriority = {
			currently_reading: 0,
			reading: 1,
			not_yet_read: 2,
			complete: 3,
		};

		return books
			.filter((book) => {
				const titleMatch = book.title?.toLowerCase().includes(query);
				const authorsMatch = book.authors
					?.toLowerCase()
					.includes(query);
				const matchesSearch = !query || titleMatch || authorsMatch;
				const matchesStatus =
					statusFilters.length === 0 ||
					statusFilters.includes(book.status?.toLowerCase());
				return matchesSearch && matchesStatus;
			})
			.sort((a, b) => {
				const aPriority = statusPriority[a.status] ?? 99;
				const bPriority = statusPriority[b.status] ?? 99;
				return aPriority - bPriority;
			});
	});

	const getBooks = async () => {
		const response = await api.get("/books/retrieve");
		books = response.map((item) => {
			return {
				...item,
				displayStatus:
					item.status === "not_yet_read"
						? "Not Yet Read"
						: item.status === "reading"
							? `${item.percent_finished * 100}% Read`
							: item.status === "complete"
								? "Completed"
								: item.status === "currently_reading"
									? `Actively Reading -- ${item.percent_finished * 100}% Read`
									: item.status,
			};
		});
	};

	const getImages = async () => {
		const response = await api.get("/images?site=books");

		images = response.map((image) => {
			return {
				...image,
				imageLink: image.image,
			};
		});
	};

	onMount(() => {
		getBooks();
		getImages();
	});

	$inspect(statusFilters);
</script>

{#snippet bookCard(item)}
	<div class="flex flex-col items-center p-8">
		<button
			class="bg-tertiary self-center mb-2 text-sm px-4 py-2 h-10 flex items-center justify-center text-center cursor-pointer"
			onclick={() => {
				if (!statusFilters.includes(item.status)) {
					statusFilters.push(item.status);
				}
			}}
		>
			Status: {item.displayStatus}
		</button>
		<div class="group">
			<div class="relative flex w-56 h-72">
				<div
					class="w-5 h-full bg-secondary/70 border-r border-black/10
					       rounded-l-sm"
				></div>

				<div
					class="relative flex flex-col items-center justify-between text-center
					       bg-secondary w-full h-full p-6 rounded-r-sm shadow-lg
					       border-l border-black/10"
				>
					<p class="text-lg font-semibold leading-tight">
						{item.title}
					</p>
					<p class="whitespace-pre-line text-sm opacity-80">
						{item.authors}
					</p>
				</div>
			</div>
		</div>
	</div>
{/snippet}

{#snippet statusFilterPill(status)}
	<div class="flex items-center gap-2 px-4 py-2 rounded-lg bg-tertiary">
		<p class="text-sm">{status}</p>
		<button
			class="shrink-0 cursor-pointer"
			onclick={() =>
				(statusFilters = statusFilters.filter((s) => s !== status))}
		>
			<X size={14} />
		</button>
	</div>
{/snippet}

{#snippet statusFilterPillbox()}
	<div class="flex flex-row items-center flex-wrap gap-2">
		<p>Status Filters:</p>
		{#each statusFilters as filter}
			{@render statusFilterPill(filter)}
		{/each}
	</div>
{/snippet}

{#snippet searchBar()}
	<label class="flex flex-col items-center gap-1 w-full max-w-xs mx-auto">
		<span>Search (book name or author name):</span>
		<input class="bg-tertiary w-full" id="search" bind:value={searchText} />
	</label>
{/snippet}

<section class="flex items-center justify-center">
	<Hero
		header={Content["books.hero.header"]}
		subtitle={Content["books.hero.subtitle"]}
	>
		<Carousel {images} />
	</Hero>
</section>
<section
	class="flex flex-col items-center justify-center gap-4 mt-8 max-w-screen mx-auto"
>
	{@render searchBar()}
	{@render statusFilterPillbox()}
</section>
<section
	class="grid justify-items-center p-12"
	style="grid-template-columns: repeat(auto-fill, minmax(14rem, 1fr));"
>
	{#each booksFiltered as book}
		{@render bookCard(book)}
	{/each}
</section>
