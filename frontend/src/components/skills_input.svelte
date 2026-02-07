<script>
  import { onMount } from 'svelte';
  import Input from "../design-system/input.svelte";
  import FormButton from "../design-system/form_button.svelte";
  import Notif from "../design-system/notif.svelte";
  import Form from "../design-system/form.svelte";
  import DataPreview from "../components/data_preview.svelte";

  import { api } from "../utils/api.svelte.js";

  let id = $state(0);
  let name = $state("");
  let category = $state("");
  let blogLink = $state("");
  let skills = $state([]);
  let editMode = $state(false);

  let loading = $state(false);
  let error = $state("");
  let success = $state("");

  let editHook = (data) => {
    if (!data) {
      id = 0;
      name = "";
      category = "";
      blogLink = "";
      editMode = false;
    }

    id = data.id;
    name = data.name;
    category = data.category;
    blogLink = data.blog_link;
    editMode = true;
  }

  let getSkills = async () => {
    const data = await api.get("/skills");

    skills = data;
  }

  let submitForm = async () => {
    const payload = {
      name,
      category,
      blog_link: blogLink
    }

    if (editMode) {
      payload.id = id
      
      try {
        await api.put("/skills", payload)
        name = ""
        category = ""
        blogLink = ""
        editMode = false

        success = "Skill been updated!"
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

  onMount(() => {
    getSkills();
  })
</script>

<div class="flex flex-row">
  <Form submitForm={submitForm} title="Add Skill" editMode={editMode} editHook={editHook}>
    <Input label="Name" bind:value={name} required={true}/>
    <Input label="Category" bind:value={category} required={true}/>
    <Input label="Blog Link" bind:value={blogLink}/>
    <FormButton loading={loading}/>
    <Notif error={error} success={success}/>
  </Form>
  <DataPreview data={skills} editHook={editHook}/>
</div>

