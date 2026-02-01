<script>
  import Input from "../design-system/input.svelte";
  import BigInput from "../design-system/big_input.svelte";
  import DateInput from "../design-system/date_input.svelte";
  import FormButton from "../design-system/form_button.svelte";
  import Notif from "../design-system/notif.svelte";
  import Form from "../design-system/form.svelte";

  import { api } from "../utils/api.svelte.js";

  let title = $state("");
  let workplace = $state("");
  let description = $state("");
  let start_date = $state("");
  let end_date = $state("");

  let loading = $state(false);
  let error = $state("");
  let success = $state("");

  async function submitForm() {
    error = "";
    success = "";
    loading = true;

    const payload = {
      title,
      workplace,
      description,
      start_date: start_date || null,
      end_date: end_date || null,
    };

    try {
      await api.post("/work-experiences", payload);

      title = "";
      workplace = "";
      description = "";
      start_date = "";
      end_date = "";

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
</script>

<Form
  submitForm={submitForm}
  title={"Add Work Experience"}
>
  <Input
    label={"Title"}
    bind:value={title}
  />

  <Input
    label={"Workplace"}
    bind:value={workplace}
  />

  <BigInput
    label={"Description"}
    bind:value={description}
  />

	<div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
    <DateInput
      label={"Start Date"}
      bind:value={start_date}
    />

    <DateInput
      label={"End Date"}
      bind:value={end_date}
    />
	</div>

  <FormButton
    loading={loading}
  />

  <Notif
    error={error}
    success={success}
  />
</Form>
