<script>
	import { ChevronRight, ChevronLeft } from "@lucide/svelte";

	let index = $state(0);

	let { images = [], intervalMs = 4000 } = $props();

	function next() {
		index = (index + 1) % images.length;
	}

	function prev() {
		index = (index - 1 + images.length) % images.length;
	}

	$effect(() => {
		if (images.length <= 1) return;

		const timer = setInterval(next, intervalMs);

		return () => clearInterval(timer);
	});
</script>

{#if images.length > 0}
	<section class="flex-1 max-w-2xl px-4 pt-4 pb-12">
		<div class="relative group">
			<div class="relative overflow-hidden aspect-[5/3]">
				<div
					class="flex h-full transition-transform duration-1000 ease-in-out"
					style="transform: translateX(-{index * 100}%);"
				>
					{#each images as image, i}
						<img
							src={image.imageLink}
							alt={image.title}
							class="w-full h-full shrink-0 object-cover"
							loading={i === 0 ? "eager" : "lazy"}
						/>
					{/each}
				</div>
			</div>

			<div class="mt-3">
				<div class="bg-secondary px-4 py-3">
					<h1 class="text-lg font-semibold">
						{images[index].title}
					</h1>
					<p class="text-sm">
						{images[index].caption}
					</p>
				</div>
			</div>
		</div>
	</section>
{/if}
