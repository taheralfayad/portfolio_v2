<script>
	let { header, subtitle, carouselImages, children } = $props();
</script>

<div class="flex flex-row p-2 justify-between gap-20 items-center text-center">
	<div class="pb-5">
		<h1 class="text-xl">{header}</h1>
		<p>
			{subtitle}
		</p>
	</div>

	{@render children()}
</div>
