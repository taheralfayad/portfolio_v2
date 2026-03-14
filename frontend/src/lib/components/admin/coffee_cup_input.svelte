<script>
  import { onMount } from "svelte";

  import Input from "$lib/design-system/input.svelte";
  import Select from "$lib/design-system/select.svelte";
  import DataPreview from "$lib/components/admin/data_preview.svelte";
  import Notif from "$lib/design-system/notif.svelte";
  import FormButton from "$lib/design-system/form_button.svelte";
  import Form from "$lib/design-system/form.svelte";
  import DateInput from "$lib/design-system/date_input.svelte";

  import { api } from "$lib/utils/api.svelte.js";

  let roasts = $state([]);
  let coffeeCups = $state([]);

  let roastId = $state(0);
  let temperature = $state(0);
  let dateDrank = $state("");
  let acidity = $state(0);
  let body = $state(0);
  let sweetness = $state(0);
  let waterType = $state("");
  let grindSize = $state(0);
  let method = $state("");
  let rating = $state(0);

  let loading = $state(false);
  let error = $state("");
  let success = $state("");

  const getCoffees = async () => {
    const data = await api.get("/coffees?include_roasts=true");

    console.log(data);

    roasts = data.flatMap((c) =>
      c.roasts.map((r) => ({
        value: r.id,
        label: `${c.name} - ${r.roast_date}`,
      })),
    );
  };

  const getCoffeeCups = async () => {
    if (!roastId) {
      coffeeCups = [];
      return;
    }
    const data = await api.get(`/coffee-cups?roast_id=${roastId}`);
    coffeeCups = data;
  };

  const submitForm = async () => {
    error = "";
    success = "";
    loading = true;

    const payload = {
      roast_id: +roastId,
      temperature: +temperature,
      date_drank: new Date(dateDrank).toISOString(),
      acidity: +acidity,
      body: +body,
      sweetness: +sweetness,
      water_type: waterType,
      grind_size: +grindSize,
      method: method,
      rating: +rating,
    };

    try {
      await api.post("/coffee-cups", payload);
      temperature = 0;
      dateDrank = "";
      acidity = 0;
      body = 0;
      sweetness = 0;
      waterType = "";
      grindSize = 0;
      method = "";
      rating = 0;
      success = "Coffee cup successfully submitted!";
      await getCoffeeCups();
    } catch (err) {
      console.error(err);
      error =
        err?.data?.error || err?.data?.message || "Failed to create coffee cup";
    } finally {
      loading = false;
    }
  };

  const handleCoffeeChange = async () => {
    await getCoffeeCups();
  };

  onMount(() => {
    getCoffees();
  });
</script>

<div class="flex flex-row">
  <Form {submitForm} title="Add Coffee Cup">
    <Select
      label="Coffee"
      bind:value={roastId}
      required={true}
      options={roasts}
      onchange={handleCoffeeChange}
    />
    <Input label="Temperature" bind:value={temperature} required={true} />
    <DateInput label="Date Drank" bind:value={dateDrank} />
    <Input label="Acidity" bind:value={acidity} required={true} />
    <Input label="Body" bind:value={body} required={true} />
    <Input label="Sweetness" bind:value={sweetness} required={true} />
    <Input label="Water Type" bind:value={waterType} required={true} />
    <Input label="Grind Size" bind:value={grindSize} required={true} />
    <Input label="Method" bind:value={method} required={true} />
    <Input label="Rating" bind:value={rating} required={true} />

    <FormButton {loading} />
    <Notif {error} {success} />
  </Form>

  <div class="ml-8">
    <DataPreview data={coffeeCups} />
  </div>
</div>
