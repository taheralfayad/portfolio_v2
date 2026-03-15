<script>
  let { title, metric } = $props();
</script>

<div
  class="relative flex flex-col items-center justify-center p-6 w-full min-w-[280px] max-w-sm mx-auto
        rounded-xl border-4 border-black
        after:content-[''] after:absolute after:inset-0
        after:rounded-xl after:bg-black
        after:translate-x-3 after:translate-y-3
        after:-z-10 bg-skills"
>
  {#if title}
    <span class="text-lg font-bold">{title}</span>
  {/if}

  {#if metric}
    <span class="text-5xl font-bold mt-2">{metric}</span>
  {/if}
</div>
