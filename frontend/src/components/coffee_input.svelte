<script>
  import { onMount } from "svelte";

  import Input from "../design-system/input.svelte";
  import BigInput from "../design-system/big_input.svelte";
  import ImageInput from "../design-system/image_input.svelte";
  import DataPreview from "./data_preview.svelte";
  import Notif from "../design-system/notif.svelte";
  import FormButton from "../design-system/form_button.svelte";
  import Form from "../design-system/form.svelte";

  import { api } from "../utils/api.svelte.js";
  import { handleImageChange } from "../utils/utils.svelte";
  import DateInput from "../design-system/date_input.svelte";

  let id = $state(0);
  let name = $state("");
  let roastLevel = $state(0);
  let roasterName = $state("");
  let originCountry = $state("");
  let processing = $state("");
  let varietal = $state("");
  let image = $state("");
  let date = $state("");
  let description = $state("");
  let coffees = $state([]);

  let editMode = $state(false);
  let loading = $state(false);
  let error = $state("");
  let success = $state("");

  const editHook = (data) => {
    if (!data) {
      id = 0;
      name = "";
      roastLevel = "";
      roasterName = "";
      originCountry = "";
      processing = "";
      varietal = "";
      image = "";
      date = "";
      description = "";
      editMode = false;
      return;
    }

    editMode = true;
    id = data.id;
    name = data.name;
    roastLevel = data.roast_level;
    roasterName = data.roaster_name;
    originCountry = data.origin_country;
    processing = data.processing;
    varietal = data.varietal;
    date = data.date;
    description = data.description;
  };

  const getCoffees = async () => {
    const data = await api.get("/coffees");

    coffees = data;
  };

  const submitForm = async () => {
    error = "";
    success = "";
    loading = true;

    const payload = {
      name: name,
      roast_level: +roastLevel,
      roaster_name: roasterName,
      origin_country: originCountry,
      processing: processing,
      varietal: varietal,
      image: image,
      date: date,
      description: description,
    };

    if (editMode) {
      payload.id = id;

      try {
        await api.put("/coffees", payload);
        name = "";
        roastLevel = 0;
        roasterName = "";
        originCountry = "";
        processing = "";
        varietal = "";
        image = "";
        date = "";
        description = "";
      } catch (err) {
        err =
          err?.data?.error || err?.data?.message || "Failed to create coffee";
      } finally {
        loading = false;
      }

      return;
    }

    try {
      await api.post("/coffees", payload);
      console.log(image);
      name = "";
      roastLevel = 0;
      roasterName = "";
      originCountry = "";
      processing = "";
      varietal = "";
      date = "";
      description = "";
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
</script>

<div class="flex flex-row">
  <Form {submitForm} {editMode} {editHook} title="Add Coffee">
    <Input label="Name" bind:value={name} required={true} />
    <Input label="Roast Level" bind:value={roastLevel} required={true} />
    <Input label="Roaster Name" bind:value={roasterName} required={true} />
    <Input label="Origin Country" bind:value={originCountry} required={true} />
    <ImageInput
      label="Image"
      onchange={async (e) =>
        (image = await handleImageChange(e.target.files[0]))}
    />
    <Input label="Processing" bind:value={processing} required={true} />
    <Input label="Varietal" bind:value={varietal} />
    <DateInput label="Roast Date" bind:value={date} />
    <BigInput label="Description" bind:value={description} />

    <FormButton {loading} />
    <Notif {error} {success} />
  </Form>
  <DataPreview data={coffees} {editHook} />
</div>
