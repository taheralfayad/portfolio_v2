<script>
  import { onMount } from "svelte";

  import Input from "$lib/design-system/input.svelte";
  import BigInput from "$lib/design-system/big_input.svelte";
  import DateInput from "$lib/design-system/date_input.svelte";
  import FormButton from "$lib/design-system/form_button.svelte";
  import Notif from "$lib/design-system/notif.svelte";
  import Form from "$lib/design-system/form.svelte";
  import DataPreview from "$lib/components/admin/data_preview.svelte";

  import { api } from "$lib/utils/api.svelte.js";
  import { normalizeDate } from "$lib/utils/utils.svelte.js";

  let id = $state(0);
  let title = $state("");
  let workplace = $state("");
  let description = $state("");
  let startDate = $state("");
  let endDate = $state("");
  let workExperiences = $state([]);
  let editMode = $state(false);

  let loading = $state(false);
  let error = $state("");
  let success = $state("");

  const editHook = (data) => {
    if (data === null) {
      editMode = false;
      id = 0;
      title = "";
      workplace = "";
      description = "";
      startDate = "";
      endDate = "";
      return;
    }
    editMode = true;
    id = data.id;
    title = data.title;
    workplace = data.workplace;
    description = data.description;
    startDate = normalizeDate(data.start_date);
    endDate = normalizeDate(data.end_date);
  };

  async function submitForm(event) {
    event.preventDefault();
    error = "";
    success = "";
    loading = true;

    const payload = {
      title,
      workplace,
      description,
      start_date: startDate || null,
      end_date: endDate || null,
    };

    if (editMode) {
      payload.id = id;
      try {
        await api.put("/work-experiences", payload);
        title = "";
        workplace = "";
        description = "";
        startDate = "";
        endDate = "";

        success = "Work Experience has been updated successfully.";
      } catch (err) {
        error =
          err?.data?.error ||
          err?.data?.message ||
          "Failed to create work experience";
      } finally {
        loading = false;
      }

      return;
    }

    try {
      await api.post("/work-experiences", payload);

      title = "";
      workplace = "";
      description = "";
      startDate = "";
      endDate = "";

      success = "Work Experience has saved successfully.";
    } catch (err) {
      error =
        err?.data?.error ||
        err?.data?.message ||
        "Failed to create work experience";
    } finally {
      loading = false;
    }
  }

  const getWorkExperiences = async () => {
    const data = await api.get("/work-experiences");

    workExperiences = data;
  };

  onMount(() => {
    getWorkExperiences();
  });
</script>

<div class="flex flex-row">
  <Form {submitForm} title={"Add Work Experience"} {editMode} {editHook}>
    <Input label={"Title"} bind:value={title} required={true} />

    <Input label={"Workplace"} bind:value={workplace} required={true} />

    <BigInput label={"Description"} bind:value={description} required={true} />

    <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
      <DateInput label={"Start Date"} bind:value={startDate} />

      <DateInput label={"End Date"} bind:value={endDate} />
    </div>

    <FormButton {loading} />

    <Notif {error} {success} />
  </Form>
  <DataPreview data={workExperiences} {editHook} />
</div>
