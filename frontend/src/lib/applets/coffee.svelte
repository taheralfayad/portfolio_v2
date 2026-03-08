<script>
  import { onMount } from "svelte";

  import DropdownTextfield from "$lib/components/home/coffee/dropdown_textfield.svelte";
  import Gauge from "$lib/components/home/coffee/gauge.svelte";
  import CoffeeDetails from "$lib/components/home/coffee/coffee_details.svelte";
  import CoffeeTable from "$lib/components/home/coffee/coffee_table.svelte";
  import Hero from "$lib/components/home/hero.svelte";
  import Carousel from "$lib/design-system/carousel.svelte";

  import { api } from "$lib/utils/api.svelte.js";
  import { formatDate } from "$lib/utils/utils.svelte";

  import coffeesNotFound from "$lib/assets/coffees_not_found.png";

  let coffees = $state([]);
  let coffeeCups = $state([]);
  let selectedCoffee = $state("");
  let searchValue = $state("");
  let isFocused = $state(false);

  let { currNavValue } = $props();

  let searchBarValues = $derived.by(() => {
    return coffees.map((x) => x.name);
  });

  let suggestions = $derived.by(() => {
    if (!searchValue) return searchBarValues;

    return searchBarValues.filter((item) =>
      item.toLowerCase().includes(searchValue.toLowerCase()),
    );
  });

  let suggestionsHidden = $derived(!isFocused);

  const getCoffees = async () => {
    const coffeesData = await api.get("/coffees");

    coffees = coffeesData.map((datum) => {
      return {
        id: datum.id,
        name: datum.name,
        roastLevel: datum.roast_level,
        roasterName: datum.roaster_name,
        originCountry: datum.origin_country,
        processing: datum.processing,
        varietal: datum.varietal,
        imageLink: datum.image,
        date: datum.date,
        description: datum.description,
      };
    });

    selectedCoffee = coffees[0];

    const coffeeCupData = await api.get(
      `/coffee-cups?coffee_id=${selectedCoffee.id}`,
    );

    coffeeCups = cleanCoffeeCupData(coffeeCupData);
  };

  const selectSuggestion = async (suggestion) => {
    selectedCoffee = coffees.find((item) => item.name === suggestion);
    isFocused = false;

    const coffeeCupData = await api.get(
      `/coffee-cups?coffee_id=${selectedCoffee.id}`,
    );

    if (coffeeCupData != null && coffeeCupData.length > 0) {
      coffeeCups = cleanCoffeeCupData(coffeeCupData);
    } else {
      coffeeCups = {};
    }
  };

  const cleanCoffeeCupData = (coffeeCupData) => {
    return coffeeCupData.map((datum) => {
      return {
        "Date Drank": formatDate(datum.date_drank),
        Temperature: datum.temperature,
        "Days After Roast": datum.days_after_roast,
        Acidity: datum.acidity,
        Body: datum.body,
        Sweetness: datum.sweetness,
        "Water Type": datum.water_type,
        "Grind Size": datum.grind_size,
        Method: datum.method,
        Rating: datum.rating,
      };
    });
  };

  onMount(() => {
    getCoffees();
  });
</script>

<section class="flex flex-col items-center justify-center">
  {#if coffees && coffees.length > 0}
    <Hero header={currNavValue.header} subtitle={currNavValue.subtitle}>
      <DropdownTextfield
        {suggestionsHidden}
        {suggestions}
        {selectSuggestion}
        currentSuggestion={selectedCoffee}
        bind:searchValue
        onFocus={() => (isFocused = true)}
      />
      <div
        class="flex flex-col md:flex-row items-center justify-center gap-12 w-full max-w-6xl mt-4"
      >
        <CoffeeDetails coffee={selectedCoffee} />
        <img
          src={selectedCoffee.imageLink}
          class="w-64 h-64 object-cover rounded-lg shadow-lg"
          alt={selectedCoffee.name}
        />
        <Gauge level={selectedCoffee.roastLevel} />
      </div>
    </Hero>
    <CoffeeTable data={coffeeCups} />
  {:else}
    <img src={coffeesNotFound} />
    <p>Sorry, I haven't added any coffees yet.</p>
  {/if}
</section>
