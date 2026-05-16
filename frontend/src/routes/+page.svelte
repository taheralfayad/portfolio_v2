<script>
  import Navbar from "$lib/components/navbar.svelte";
  import { onMount } from "svelte";
  import Home from "$lib/applets/home.svelte";
  import Coffee from "$lib/applets/coffee.svelte";
  import { page } from "$app/stores";
  import { goto } from "$app/navigation";

  const navItems = [
    {
      value: "Home",
      header: "Hello! My name is Taher Alfayad.",
      subtitle:
        "I'm a husband, older brother, dog dad, Arsenal fan, and software developer.",
      onClick: () => {
        currNavValue = navItems.find((item) => item.value === "Home");
        goto("?tab=Home", { replaceState: true });
      },
    },
    {
      value: "Coffee",
      header: "Hello! Join me on my coffee journey.",
      subtitle:
        "I love drinking coffee. I've recently started roasting my own, and this is where I document it.",
      onClick: () => {
        currNavValue = navItems.find((item) => item.value === "Coffee");
        goto("?tab=Coffee", { replaceState: true });
      },
    },
  ];

  let currNavValue = $state(navItems.find((item) => item.value === "Home"));

  onMount(() => {
    const tab = $page.url.searchParams.get("tab");
    const match = navItems.find(
      (item) => item.value.toLowerCase() === tab.toLowerCase(),
    );
    if (match) currNavValue = match;
  });

  const getCurrNavValue = () => currNavValue.value;
</script>

<Navbar items={navItems} {getCurrNavValue} />
<div class="flex flex-col min-h-screen pt-20">
  {#if currNavValue.value === "Home"}
    <Home {currNavValue} />
  {:else if currNavValue.value === "Coffee"}
    <Coffee {currNavValue} />
  {/if}
</div>
