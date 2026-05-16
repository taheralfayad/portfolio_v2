<script>
  import { onMount } from "svelte";

  import DropdownTextfield from "$lib/components/home/coffee/dropdown_textfield.svelte";
  import Gauge from "$lib/components/home/coffee/gauge.svelte";
  import CoffeeDetails from "$lib/components/home/coffee/coffee_details.svelte";
  import ErrorCard from "$lib/components/home/coffee/error_card.svelte";
  import CoffeeSql from "$lib/components/home/coffee/coffee_sql.svelte";
  import CoffeeCharts from "$lib/components/home/coffee/coffee_charts.svelte";
  import Hero from "$lib/components/home/hero.svelte";
  import Select from "$lib/design-system/select.svelte";

  import { api } from "$lib/utils/api.svelte.js";
  import { formatDate } from "$lib/utils/utils.svelte";

  import coffeesNotFound from "$lib/assets/coffees_not_found.png";

  import { ZipReader, BlobReader, BlobWriter } from "@zip.js/zip.js";

  let coffees = $state([]);
  let coffeeCups = $state([]);
  let selectedCoffee = $state({});
  let selectedRoast = $state({});
  let graphData = $state({});
  let searchValue = $state("");
  let isFocused = $state(false);
  let parquetBuffers = $state();
  let selectedView = $state("CoffeeCharts");

  let { currNavValue } = $props();

  let roasts = $derived(
    selectedCoffee.roasts
      ? [...selectedCoffee.roasts].sort(
          (a, b) => new Date(b.roast_date) - new Date(a.roast_date),
        )
      : [],
  );

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
    const coffeesData = await api.get("/coffees?include_roasts=true");

    coffees = coffeesData.map((datum) => {
      return {
        id: datum.id,
        name: datum.name,
        originCountry: datum.origin_country,
        processing: datum.processing,
        varietal: datum.varietal,
        description: datum.description,
        roasts: datum.roasts,
      };
    });

    selectedCoffee = coffees[0];
    selectedRoast = roasts.length > 0 ? roasts[0] : {};
  };

  const getDuckDBFile = async () => {
    const response = await api.download("/duckdbify");

    const zipReader = new ZipReader(new BlobReader(response));
    const entries = await zipReader.getEntries();

    const buffers = {};
    for (const entry of entries) {
      const blob = await entry.getData(new BlobWriter());
      const buffer = await blob.arrayBuffer();
      buffers[entry.filename] = new Uint8Array(buffer);
    }
    await zipReader.close();

    parquetBuffers = buffers;
  };

  const getCoffeeCups = async (roast) => {
    const roastId = roast?.id ?? selectedRoast.id;
    const coffeeCupData = await api.get(`/coffee-cups?roast_id=${roastId}`);

    coffeeCups = cleanCoffeeCupData(coffeeCupData);
  };

  const selectSuggestion = async (suggestion) => {
    selectedCoffee = coffees.find((item) => item.name === suggestion);
    selectedRoast = roasts.length > 0 ? roasts[0] : {};
    isFocused = false;
  };

  let initialLoad = true;

  $effect(() => {
    if (selectedRoast?.id) {
      if (initialLoad) {
        initialLoad = false;
        getCoffeeCups();
      } else {
        getCoffeeCups(selectedRoast);
      }
    }
  });

  const cleanCoffeeCupData = (coffeeCupData) => {
    return coffeeCupData
      ? coffeeCupData.map((datum) => {
          return {
            "Date Drank": formatDate(datum.date_drank),
            "Date Drank Raw": new Date(datum.date_drank),
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
        })
      : [];
  };

  onMount(() => {
    getCoffees();
    getDuckDBFile();
  });
</script>

<section class="flex flex-col items-center justify-center">
  {#if coffees && coffees.length > 0}
    <div class="mt-6 w-full">
      <ul class="hidden sm:flex flex-row mb-0 justify-center gap-52 px-6">
        <li>
          <button
            class="
            py-2 px-16 cursor-pointer
            relative
            border-4 {selectedView === 'CoffeeCharts'
              ? 'bg-skills font-bold'
              : ''}
            hover:bg-skills/50
            "
            onclick={() => {
              selectedView = "CoffeeCharts";
            }}>CoffeeCharts</button
          >
        </li>
        <li>
          <button
            class="
            py-2 px-16 cursor-pointer
            relative
            border-4 {selectedView === 'CoffeeSQL' ? 'bg-skills font-bold' : ''}
            hover:bg-skills/50
            "
            onclick={() => {
              selectedView = "CoffeeSQL";
            }}>CoffeeSQL</button
          >
        </li>
      </ul>
      <div class="sm:p-6 flex flex-col">
        {#if selectedView === "CoffeeCharts" && roasts && roasts.length > 0}
          {#if coffeeCups && coffeeCups.length > 0}
            <Hero header={currNavValue.header} subtitle={currNavValue.subtitle}>
              <div class="flex flex-col sm:flex-row gap-6">
                <DropdownTextfield
                  {suggestionsHidden}
                  {suggestions}
                  {selectSuggestion}
                  currentSuggestion={selectedCoffee}
                  bind:searchValue
                  onFocus={() => (isFocused = true)}
                  onClickOutside={() => (isFocused = false)}
                />
                {#if roasts}
                  <Select
                    label="Then, get the roast:"
                    bind:value={selectedRoast}
                    options={roasts
                      ? roasts.map((r) => {
                          return { value: r, label: formatDate(r.roast_date) };
                        })
                      : { value: null, label: "No Roasts Found" }}
                  />
                {/if}
              </div>
              <div
                class="flex flex-col md:flex-row items-center justify-center gap-24 w-full max-w-6xl mt-8"
              >
                <CoffeeDetails
                  name={selectedCoffee.name}
                  originCountry={selectedCoffee.originCountry}
                  processing={selectedCoffee.processing}
                  varietal={selectedCoffee.varietal}
                  date={selectedRoast.roast_date}
                  description={selectedCoffee.description}
                />
                {#if roasts && roasts.length > 0}
                  <img
                    src={selectedRoast.image}
                    class="max-w-96 max-h-80 object-cover"
                    alt={selectedCoffee.name + "-" + selectedRoast.roast_date}
                  />
                  <Gauge level={selectedRoast.roast_level} />
                {:else if roasts && roasts.length === 0}
                  <ErrorCard message="No Roasts Found" />
                {/if}
              </div>
            </Hero>
            <CoffeeCharts {coffeeCups} />
          {:else}
            <div>
              <ErrorCard message="No Coffee Cups Found" />
            </div>
          {/if}
        {:else if selectedView === "CoffeeSQL"}
          <div class="flex flex-col items-center justify-center">
            <CoffeeSql
              parquetBuffers={parquetBuffers
                ? Object.fromEntries(
                    Object.entries(parquetBuffers).map(([k, v]) => [
                      k,
                      v.slice(0),
                    ]),
                  )
                : undefined}
            />
          </div>
        {/if}
      </div>
    </div>
  {:else}
    <img src={coffeesNotFound} />
    <p>Sorry, I haven't added any coffees yet.</p>
  {/if}
</section>
