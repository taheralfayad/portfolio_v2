<script>
  let { coffee } = $props();

  import { formatDate } from "$lib/utils/utils.svelte";
</script>

<div
  class="relative flex flex-col items-center p-6 w-full min-w-[280px] max-w-sm mx-auto
        rounded-xl border-4 border-black
        after:content-[''] after:absolute after:inset-0
        after:rounded-xl after:bg-black
        after:translate-x-3 after:translate-y-3
        after:-z-10 bg-skills"
>
  {#if coffee.originCountry}
    <div class="w-full">
      <span class="font-bold">Name:</span>
      <span>{coffee.name}</span>
    </div>
  {/if}

  {#if coffee.originCountry}
    <div class="w-full mt-3">
      <span class="font-bold">Origin:</span>
      <span>{coffee.originCountry}</span>
    </div>
  {/if}

  {#if coffee.processing}
    <div class="w-full mt-3">
      <span class="font-bold">Processing:</span>
      <span>{coffee.processing}</span>
    </div>
  {/if}

  {#if coffee.varietal}
    <div class="w-full mt-3">
      <span class="font-bold">Varietal:</span>
      <span>{coffee.varietal}</span>
    </div>
  {/if}

  {#if coffee.date}
    <div class="w-full mt-3">
      <span class="font-bold">Roast Date:</span>
      <span>{formatDate(coffee.date)}</span>
    </div>
  {/if}

  {#if coffee.description}
    <div class="w-full mt-3">
      <span class="font-bold">Notes:</span>
      <p class="mt-1">{coffee.description}</p>
    </div>
  {/if}
</div>
