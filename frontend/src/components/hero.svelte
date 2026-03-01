<script>
  import Carousel from "../design-system/carousel.svelte";

  let { header, subtitle, carouselImages } = $props();
</script>

<div class="flex flex-col p-2 mt-12 items-center text-center">
  <h1 class="text-xl">{header}</h1>
  <p>
    {subtitle}
  </p>
  <Carousel images={carouselImages} />
</div>
