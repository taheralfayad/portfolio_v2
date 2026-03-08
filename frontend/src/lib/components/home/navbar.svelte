<script>
  let { items, getCurrNavValue} = $props();

  import NavItem from "../design-system/nav_item.svelte";
</script>

<nav class="absolute top-0 left-0 right-0 z-50 bg-transparent">
  <ul class="flex justify-center gap-8 pt-6 pb-4">
    {#each items as item}
      <NavItem
        value={item.value}
        onClick={item.onClick}
        getCurrNavValue={getCurrNavValue}
      />
    {/each}
  </ul>
</nav>
