<script>
  import { onMount } from "svelte";

  import DropdownTextfield from "$lib/components/home/coffee/dropdown_textfield.svelte";
  import Gauge from "$lib/components/home/coffee/gauge.svelte";
  import CoffeeDetails from "$lib/components/home/coffee/coffee_details.svelte";
  import CoffeeKpiCard from "$lib/components/home/coffee/coffee_kpi_card.svelte";
  import CoffeeTable from "$lib/components/home/coffee/coffee_table.svelte";
  import ErrorCard from "$lib/components/home/coffee/error_card.svelte";
  import Hero from "$lib/components/home/hero.svelte";
  import Carousel from "$lib/design-system/carousel.svelte";
  import Select from "$lib/design-system/select.svelte";
  import { chartRender } from "$lib/actions/chartRender.svelte";

  import { api } from "$lib/utils/api.svelte.js";
  import { formatDate } from "$lib/utils/utils.svelte";

  import coffeesNotFound from "$lib/assets/coffees_not_found.png";

  let coffees = $state([]);
  let coffeeCups = $state([]);
  let selectedCoffee = $state({});
  let selectedRoast = $state({});
  let graphData = $state({});
  let selectedLabel = $state("Date Drank");
  let selectedMetric = $state("Rating");
  let selectedFilter = $state("");
  let selectedFilterValue = $state("");
  let searchValue = $state("");
  let isFocused = $state(false);

  let { currNavValue } = $props();

  let roasts = $derived(selectedCoffee.roasts);

  let average = $derived.by(() => {
    if (
      !selectedMetric ||
      !filteredCoffeeCups ||
      filteredCoffeeCups.length === 0
    ) {
      return;
    }
    const values = filteredCoffeeCups.map((c) => {
      return c[selectedMetric];
    });

    const sum = values.reduce(
      (accumulator, currentValue) => accumulator + currentValue,
      0,
    );

    const avg = sum / values.length;

    return avg.toFixed(1);
  });

  let min = $derived(
    filteredCoffeeCups && filteredCoffeeCups.length > 0
      ? Math.min(...filteredCoffeeCups.map((c) => c[selectedMetric]))
      : undefined,
  );

  let max = $derived(
    filteredCoffeeCups && filteredCoffeeCups.length > 0
      ? Math.max(...filteredCoffeeCups.map((c) => c[selectedMetric]))
      : undefined,
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

  let selectedFilterValues = $derived.by(() => {
    return [...new Set(coffeeCups.map((c) => c[selectedFilter]))];
  });

  let filteredCoffeeCups = $derived.by(() => {
    let filtered = [...coffeeCups];

    if (selectedFilter && selectedFilterValue) {
      filtered = filtered.filter(
        (c) => c[selectedFilter] === selectedFilterValue,
      );
    }

    return filtered;
  });

  let coffeeData = $derived.by(() => {
    let filtered = [...coffeeCups];

    if (selectedFilter && selectedFilterValue) {
      filtered = filtered.filter(
        (c) => c[selectedFilter] === selectedFilterValue,
      );
    }

    const sorted = filtered.sort((a, b) => {
      const sortKey =
        selectedLabel === "Date Drank" ? "Date Drank Raw" : selectedLabel;
      const labelA = a[sortKey];
      const labelB = b[sortKey];

      if (labelA instanceof Date) return labelA.getTime() - labelB.getTime();
      if (typeof labelA === "string") return labelA.localeCompare(labelB);

      return labelA - labelB;
    });

    return {
      type: "line",
      data: {
        labels: sorted.map((c) => c[selectedLabel]),
        datasets: [
          {
            label: selectedMetric,
            data: sorted.map((c) => c[selectedMetric]),
            backgroundColor: ["rgba(212, 202, 189)"],
          },
        ],
      },
    };
  });

  const LABELS = ["Date Drank", "Temperature"];

  const METRICS = ["Rating", "Acidity", "Sweetness", "Body"];

  const FILTERS = ["", "Method", "Water Type"];

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
    selectedRoast =
      selectedCoffee.roasts && selectedCoffee.roasts.length > 0
        ? selectedCoffee.roasts[0]
        : {};
  };

  const getCoffeeCups = async (roast) => {
    const roastId = roast?.id ?? selectedRoast.id;
    const coffeeCupData = await api.get(`/coffee-cups?roast_id=${roastId}`);

    coffeeCups = cleanCoffeeCupData(coffeeCupData);
  };

  const selectSuggestion = async (suggestion) => {
    selectedCoffee = coffees.find((item) => item.name === suggestion);
    selectedRoast =
      selectedCoffee.roasts && selectedCoffee.roasts.length > 0
        ? selectedCoffee.roasts[0]
        : {};
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
  });
</script>

<section class="flex flex-col items-center justify-center">
  {#if coffees && coffees.length > 0}
    <Hero header={currNavValue.header} subtitle={currNavValue.subtitle}>
      <div class="flex flex-col sm:flex-row gap-6 overflow-x-scroll">
        <DropdownTextfield
          {suggestionsHidden}
          {suggestions}
          {selectSuggestion}
          currentSuggestion={selectedCoffee}
          bind:searchValue
          onFocus={() => (isFocused = true)}
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
        class="flex flex-col md:flex-row items-center justify-center gap-12 w-full max-w-6xl mt-4"
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
            class="w-64 h-64 object-cover rounded-lg shadow-lg"
            alt={selectedCoffee.name + "-" + selectedRoast.roast_date}
          />
          <Gauge level={selectedRoast.roast_level} />
        {:else if roasts && roasts.length === 0}
          <ErrorCard message="No Roasts Found" />
        {/if}
      </div>
    </Hero>
    {#if roasts && roasts.length > 0}
      <div class="flex flex-col sm:flex-row gap-4 my-12">
        <CoffeeKpiCard title={`Average ${selectedMetric}`} metric={average} />
        <CoffeeKpiCard title={`Minimum ${selectedMetric}`} metric={min} />
        <CoffeeKpiCard title={`Maximum ${selectedMetric}`} metric={max} />
      </div>
      {#if coffeeCups && coffeeCups.length > 0}
        <div
          class="flex flex-col gap-4 w-full items-center overflow-scroll mx-auto"
        >
          <div
            class="flex flex-col sm:flex-row sm:flex-wrap gap-2 sm:gap-4 sm:justify-center"
          >
            <Select
              label="Labels"
              bind:value={selectedLabel}
              options={LABELS.map((l) => {
                return { value: l, label: l };
              })}
            />
            <Select
              label="Metric"
              bind:value={selectedMetric}
              options={METRICS.map((m) => {
                return { value: m, label: m };
              })}
            />
            <Select
              label="Filter"
              bind:value={selectedFilter}
              options={FILTERS.map((f) => {
                return { value: f, label: f };
              })}
            />
            {#if selectedFilter}
              <Select
                label="Filter Value"
                bind:value={selectedFilterValue}
                options={selectedFilterValues.map((v) => {
                  return { value: v, label: v };
                })}
              />
            {/if}
          </div>
        </div>
        <div class="h-64">
          <canvas use:chartRender={coffeeData} class="md:w-200"></canvas>
        </div>
        <CoffeeTable
          data={coffeeCups.map((c) =>
            Object.fromEntries(
              Object.entries(c).filter(([k]) => k !== "Date Drank Raw"),
            ),
          )}
        />
      {:else}
        <div>
          <ErrorCard message="No Coffee Cups Found" />
        </div>
      {/if}
    {/if}
  {:else}
    <img src={coffeesNotFound} />
    <p>Sorry, I haven't added any coffees yet.</p>
  {/if}
</section>
