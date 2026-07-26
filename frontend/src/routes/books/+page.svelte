<script>
	import { onMount } from "svelte";
	import { api } from "$lib/utils/api.svelte";

	import Hero from "$lib/components/home/hero.svelte";
	import Carousel from "$lib/design-system/carousel.svelte";

	import Content from "$lib/content/home.json";

	let books = $state([]);
	let images = $state([]);

	const getBooks = async () => {
		const response = await api.get("/get-books");
		books = response.map((item) => {
			return {
				...item,
				status:
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
</script>

{#snippet bookCard(item)}
	<div class="flex flex-col items-center p-8">
		<div
			class="bg-tertiary self-center mb-2 text-sm px-4 py-2 h-10 flex items-center justify-center text-center"
		>
			Status: {item.status}
		</div>
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

<section class="flex items-center justify-center">
	<Hero
		header={Content["books.hero.header"]}
		subtitle={Content["books.hero.subtitle"]}
	>
		<Carousel {images} />
	</Hero>
</section>
<section
	class="grid mt-8 justify-items-center p-12"
	style="grid-template-columns: repeat(auto-fill, minmax(14rem, 1fr));"
>
	{#each books as book}
		{@render bookCard(book)}
	{/each}
</section>
