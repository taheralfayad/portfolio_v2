<script>
  let index = $state(0);

  let { images = [] } = $props();

  function next() {
    index = (index + 1) % images.length;
  }

  function prev() {
    index = (index - 1 + images.length) % images.length;
  }
</script>

{#if images.length > 0}
<section class="w-full max-w-2xl mx-auto px-4 pt-4 pb-12">
  <div class="relative group">
    <div class="relative overflow-hidden rounded-2xl shadow-lg aspect-[5/3] bg-neutral-100">
      <div
        class="flex h-full transition-transform duration-1000 ease-in-out"
        style="transform: translateX(-{index * 100}%);"
      >
        {#each images as image, i}
          <img
            src={image.imageLink}
            alt={image.title}
            class="w-full h-full flex-shrink-0 object-cover"
            loading={i === 0 ? 'eager' : 'lazy'}
          />
        {/each}
      </div>

      <div class="absolute inset-0 bg-gradient-to-t from-black/30 via-black/0 to-black/0" />

      <div class="absolute bottom-0 left-0 right-0 p-6">
        <div class="inline-block rounded-xl bg-white/80 backdrop-blur px-4 py-3 shadow-sm">
          <h1 class="text-lg font-semibold text-gray-900">
            {images[index].title}
          </h1>
          <p class="text-sm text-gray-600">
            {images[index].caption}
          </p>
        </div>
      </div>

      <button
        onclick={prev}
        aria-label="Previous slide"
        class="absolute left-3 top-1/2 -translate-y-1/2
               w-8 h-8 rounded-full
               bg-white/70 backdrop-blur
               text-gray-700
               shadow-sm
               opacity-0 group-hover:opacity-100
               transition hover:bg-white
               hover:cursor-pointer"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="mx-auto" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m15 18-6-6 6-6" />
        </svg>
      </button>

      <button
        onclick={next}
        aria-label="Next slide"
        class="absolute right-3 top-1/2 -translate-y-1/2
               w-8 h-8 rounded-full
               bg-white/70 backdrop-blur
               text-gray-700
               shadow-sm
               opacity-0 group-hover:opacity-100
               transition hover:bg-white
               hover:cursor-pointer"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="mx-auto" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m9 18 6-6-6-6" />
        </svg>
      </button>
    </div>

  </div>
</section>
{/if}
