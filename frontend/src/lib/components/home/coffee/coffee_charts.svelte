<script>
  import Select from "$lib/design-system/select.svelte";
  import CoffeeKpiCard from "$lib/components/home/coffee/coffee_kpi_card.svelte";
  import CoffeeTable from "$lib/components/home/coffee/coffee_table.svelte";

  import { chartRender } from "$lib/actions/chartRender.svelte";

  let {
    average,
    min,
    max,
    selectedFilterValues,
    coffeeCups,
    selectedLabel = $bindable(),
    selectedMetric = $bindable(),
    selectedFilter = $bindable(),
    selectedFilterValue = $bindable(),
  } = $props();

  const LABELS = ["Date Drank", "Temperature"];
  const METRICS = ["Rating", "Acidity", "Sweetness", "Body"];
  const FILTERS = ["", "Method", "Water Type"];

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
            borderColor: "#000000",
            pointBackgroundColor: "#000000",
            pointBorderColor: "#000000",
          },
        ],
      },
    };
  });
</script>

<div
  class="flex flex-col gap-4 w-full items-center overflow-auto mx-auto my-12"
>
  <div
    class="flex flex-col justify-center items-center sm:flex-row sm:flex-wrap gap-2 sm:gap-4 sm:justify-center"
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
<div class="flex flex-col sm:flex-row gap-4 mb-6">
  <CoffeeKpiCard title={`Average ${selectedMetric}`} metric={average} />
  <CoffeeKpiCard title={`Minimum ${selectedMetric}`} metric={min} />
  <CoffeeKpiCard title={`Maximum ${selectedMetric}`} metric={max} />
</div>
<div class="h-64 w-full max-w-5xl items-center justify-center mx-auto">
  <canvas
    class="border border-black text-black text-md"
    use:chartRender={coffeeData}
  ></canvas>
</div>
<CoffeeTable
  data={coffeeCups.map((c) =>
    Object.fromEntries(
      Object.entries(c).filter(([k]) => k !== "Date Drank Raw"),
    ),
  )}
/>
