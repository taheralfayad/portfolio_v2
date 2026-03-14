<script>
  let { message } = $props();
</script>

<div
  class="relative flex items-center justify-center p-8 w-full max-w-md mx-auto
        rounded-xl border-4 border-black
        after:content-[''] after:absolute after:inset-0
        after:rounded-xl after:bg-black
        after:translate-x-3 after:translate-y-3
        after:-z-10 bg-red-200"
>
  <p class="font-bold text-lg text-center">{message}</p>
</div>
