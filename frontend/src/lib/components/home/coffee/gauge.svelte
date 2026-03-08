<script>
  let { level } = $props();

  let cityFilled = "bg-[#F5DEB3]";
</script>

<div class="flex flex-row items-center gap-4">
  <div class="flex flex-col">
    <div class="flex flex-col border-4 border-black rounded-md overflow-hidden">
      <div
        class="w-12 h-12 border-t-4 border-black flex items-center justify-center
              {level >= 5 ? 'bg-[#1C0A05]' : 'bg-white'}"
      >
        {#if level === 5}
          <span
            class="text-[12px] font-bold text-white uppercase tracking-tight leading-none text-center"
            >French</span
          >
        {/if}
      </div>
      <div
        class="w-12 h-12 border-t-4 border-black flex items-center justify-center
              {level >= 4 ? 'bg-[#4A2C1A]' : 'bg-white'}"
      >
        {#if level === 4}
          <span
            class="text-[12px] font-bold text-white uppercase tracking-tight leading-none text-center"
            >Full<br />City+</span
          >{/if}
      </div>
      <div
        class="w-12 h-12 border-t-4 border-black flex items-center justify-center
              {level >= 3 ? 'bg-[#8B5E3C]' : 'bg-white'}"
      >
        {#if level === 3}
          <span
            class="text-[12px] font-bold text-white uppercase tracking-tight leading-none text-center"
            >Full<br />City</span
          >
        {/if}
      </div>
      <div
        class="w-12 h-12 border-t-4 border-black flex items-center justify-center
              {level >= 2 ? 'bg-[#C8924A]' : 'bg-white'}"
      >
        {#if level === 2}
          <span
            class="text-[12px] font-bold text-black uppercase tracking-tight leading-none text-center"
            >City+</span
          >
        {/if}
      </div>
      <div
        class="w-12 h-12 border-t-4 border-black flex items-center justify-center
              {level >= 1 ? 'bg-[#F5DEB3]' : 'bg-white'}"
      >
        {#if level === 1}
          <span
            class="text-[12px] font-bold text-black uppercase tracking-tight leading-none text-center"
            >City</span
          >
        {/if}
      </div>
    </div>
  </div>
  <h2 class="text-4xl font-black uppercase leading-none tracking-tight">
    Roast<br />Level
  </h2>
</div>
