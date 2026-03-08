<script>
  import Carousel from "$lib/design-system/carousel.svelte";

  let { header, subtitle, carouselImages, children } = $props();
</script>

<div class="flex flex-col p-2 items-center text-center">
  <div class="pb-5">
    <h1 class="text-xl">{header}</h1>
    <p>
      {subtitle}
    </p>
  </div>

  {@render children()}

</div>
