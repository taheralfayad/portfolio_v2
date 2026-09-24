<script>
	import { onMount } from "svelte";

	import { api } from "$lib/utils/api.svelte";

	import Hero from "$lib/components/home/hero.svelte";
	import Carousel from "$lib/design-system/carousel.svelte";
	import FilterBar from "$lib/components/home/filter_bar.svelte";
	import LoadingSpinner from "$lib/components/home/books/loading_spinner.svelte";

	import Content from "$lib/content/home.json";

	let books = $state([]);
	let images = $state([]);
	let searchText = $state("");
	let statusFilters = $state([]);
	let isLoading = $state(true);

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
							? `${Math.round(item.percent_finished * 100)}% Read`
							: item.status === "complete"
								? "Completed"
								: item.status === "currently_reading"
									? `Actively Reading -- ${Math.round(item.percent_finished * 100)}% Read`
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

	onMount(async () => {
		try {
			await Promise.all([getBooks(), getImages()]);
		} catch (err) {
			console.error(err);
		} finally {
			isLoading = false;
		}
	});

	const statusOptions = [
		{ value: "currently_reading", label: "Actively Reading" },
		{ value: "reading", label: "Reading" },
		{ value: "not_yet_read", label: "Not Yet Read" },
		{ value: "complete", label: "Completed" },
	];
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

{#if isLoading}
	<LoadingSpinner />
{:else}
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
		<FilterBar
			bind:searchText
			bind:selectedOptions={statusFilters}
			options={statusOptions}
			searchPlaceholder="Search (book name or author name)..."
			pillboxLabel="Status Filters"
		/>
	</section>
	<section
		class="grid justify-items-center p-12"
		style="grid-template-columns: repeat(auto-fill, minmax(14rem, 1fr));"
	>
		{#each booksFiltered as book}
			{@render bookCard(book)}
		{/each}
	</section>
{/if}
