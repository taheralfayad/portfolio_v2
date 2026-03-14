<script>
  import { onMount } from "svelte";

  import Input from "$lib/design-system/input.svelte";
  import DateInput from "$lib/design-system/date_input.svelte";
  import ImageInput from "$lib/design-system/image_input.svelte";
  import Form from "$lib/design-system/form.svelte";
  import FormButton from "$lib/design-system/form_button.svelte";
  import Notif from "$lib/design-system/notif.svelte";
  import Select from "$lib/design-system/select.svelte";

  import DataPreview from "$lib/components/admin/data_preview.svelte";

  import { api } from "$lib/utils/api.svelte";
  import { handleImageChange } from "$lib/utils/utils.svelte";

  let coffees = $state([]);
  let coffeeId = $state(0);

  let id = $state(0);
  let roastLevel = $state(0);
  let roasterName = $state("");
  let roastDate = $state("");
  let image = $state("");

  let editMode = $state(false);
  let loading = $state(false);
  let error = $state("");
  let success = $state("");

  let roasts = $derived.by(() => {
    return coffees.find((item) => item.value === coffeeId)?.roasts;
  });

  const getCoffees = async () => {
    const data = await api.get("/coffees?include_roasts=true");
    coffees = data.map((c) => ({
      value: c.id,
      label: c.name,
      roasts: c.roasts,
    }));

    console.log(coffees);
  };

  const submitForm = async () => {
    error = "";
    success = "";
    loading = true;

    const payload = {
      roast_level: +roastLevel,
      roaster_name: roasterName,
      roast_date: roastDate,
      image: image,
      coffee_id: +coffeeId,
    };

    try {
      await api.post("/roasts", payload);
      roastLevel = 0;
      roasterName = "";
      roastDate = "";
      image = "";
    } catch (err) {
      console.error(err);
      error = err || "Failed to create coffee";
    } finally {
      loading = false;
    }
    return;
  };

  onMount(() => {
    getCoffees();
  });

  $inspect(roasts);
</script>

<div class="flex flex-row">
  <Form {submitForm} title="Add Roast">
    <Select
      label="Coffee"
      bind:value={coffeeId}
      required={true}
      options={coffees}
    />
    <Input label="Roast Level" bind:value={roastLevel} required={true} />
    <Input label="Roaster Name" bind:value={roasterName} required={true} />
    <DateInput label="Roast Date" bind:value={roastDate} />
    <ImageInput
      label="Image"
      onchange={async (e) =>
        (image = await handleImageChange(e.target.files[0]))}
    />
    <FormButton {loading} />
    <Notif {error} {success} />
  </Form>
  <DataPreview data={roasts} />
</div>
