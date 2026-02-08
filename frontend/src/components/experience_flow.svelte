<script>
  export let items = [];

</script>

<div class="flex flex-wrap items-center justify-center gap-6 py-12 px-6 max-w-6xl mx-auto"> 
  {#each items as item, index}
    <div class="relative flex-shrink-0">
      <div
        class="w-64 min-h-[280px] rounded-xl border-3 border-black bg-experience p-6 text-center flex flex-col items-center justify-center
          after:content-[''] after:absolute after:inset-0
          after:rounded-xl after:bg-black
          after:translate-x-2 after:translate-y-2
          after:-z-10
        "
      >
        <div class="text-base font-bold leading-tight">
          {item.title}
        </div>
        {#if item.subtitle}
          <div class="mt-2 text-sm">
            {item.subtitle}
          </div>
        {/if}
        {#if item.description}
          <div class="mt-3 text-xs leading-relaxed">
            {item.description}
          </div>
        {/if}
      </div>
    </div>
    {#if index < items.length - 1}
      <div class="hidden sm:block select-none text-2xl shrink-0">
        →
      </div>
    {/if}
  {/each}
</div>
