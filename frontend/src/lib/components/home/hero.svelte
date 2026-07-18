<script>
	let { header, subtitle, carouselImages, children } = $props();
</script>

<div
	class="flex flex-wrap sm:flex-row p-2 sm:justify-between md:gap-20 md:mt-6 items-center text-center"
>
	<div class="pb-5 max-sm:mt-4">
		<h1 class="text-xl">{header}</h1>
		<p>
			{subtitle}
		</p>
	</div>

	{@render children()}
</div>
