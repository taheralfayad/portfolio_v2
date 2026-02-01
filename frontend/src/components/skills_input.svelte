<script>
  import Input from "../design-system/input.svelte";
  import FormButton from "../design-system/form_button.svelte";
  import Notif from "../design-system/notif.svelte";
  import Form from "../design-system/form.svelte";

  import { api } from "../utils/api.svelte.js";

  let name = $state("");
  let category = $state("");
  let blogLink = $state("");

  let loading = $state(false);
  let error = $state("");
  let success = $state("");

  let submitForm = async () => {
    const payload = {
      name,
      category,
      blog_link: blogLink
    }

    try {
      await api.post("/skills", payload)
      success = "Skill uploaded successfully!"
    } catch (err) {
      console.error(err)
      error = 
        err || 
        "Failed to create skill";
    } finally {
      loading = false
    }
  }
</script>

<Form submitForm={submitForm} title="Add Skill">
  <Input label="Name" bind:value={name}/>
  <Input label="Category" bind:value={category}/>
  <Input label="Blog Link" bind:value={blogLink}/>
  <FormButton loading={loading}/>
  <Notif error={error} success={success}/>
</Form>
