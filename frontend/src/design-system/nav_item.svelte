<script>
  let { value, onClick, getCurrNavValue } = $props();
</script>
<li
  class="hover:bg-button/80 transition p-3 px-6 rounded-md cursor-pointer"
  onclick={onClick}
>
  <span class="text-sm {getCurrNavValue() === value ? 'font-black' : 'font-mono'}">{value}</span>
</li>
