<script>
  let { name, originCountry, processing, varietal, date, description } =
    $props();

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
  {#if name}
    <div class="w-full">
      <span class="font-bold">Name:</span>
      <span>{name}</span>
    </div>
  {/if}

  {#if originCountry}
    <div class="w-full mt-3">
      <span class="font-bold">Origin:</span>
      <span>{originCountry}</span>
    </div>
  {/if}

  {#if processing}
    <div class="w-full mt-3">
      <span class="font-bold">Processing:</span>
      <span>{processing}</span>
    </div>
  {/if}

  {#if varietal}
    <div class="w-full mt-3">
      <span class="font-bold">Varietal:</span>
      <span>{varietal}</span>
    </div>
  {/if}

  {#if date}
    <div class="w-full mt-3">
      <span class="font-bold">Roast Date:</span>
      <span>{formatDate(date)}</span>
    </div>
  {/if}

  {#if description}
    <div class="w-full mt-3">
      <span class="font-bold">Notes:</span>
      <p class="mt-1">{description}</p>
    </div>
  {/if}
</div>
