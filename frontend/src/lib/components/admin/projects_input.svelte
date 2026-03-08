<script>
  import { onMount } from "svelte";

  import Input from "$lib/design-system/input.svelte";
  import BigInput from "$lib/design-system/big_input.svelte";
  import DateInput from "$lib/design-system/date_input.svelte";
  import ImageInput from "$lib/design-system/image_input.svelte";
  import FormButton from "$lib/design-system/form_button.svelte";
  import Notif from "$lib/design-system/notif.svelte";
  import Form from "$lib/design-system/form.svelte";
  import DataPreview from "$lib/components/admin/data_preview.svelte";

  import { api } from "$lib/utils/api.svelte.js";
  import { handleImageChange } from "$lib/utils/utils.svelte";

  let id = $state(0);
  let name = $state("");
  let description = $state("");
  let githubLink = $state("");
  let image = $state("");
  let blogLink = $state("");
  let type = $state("");
  let projects = $state([]);
  let editMode = $state(false);

  let error = $state("");
  let success = $state("");
  let loading = $state(false);

  const editHook = (data) => {
    if (!data) {
      name = "";
      description = "";
      githubLink = "";
      image = "";
      blogLink = "";
      type = "";
    }

    id = data.id;
    name = data.name;
    description = data.description;
    githubLink = data.github_link;
    blogLink = data.blog_link;
    type = data.type;
    editMode = true;
  };

  const getProjects = async () => {
    const data = await api.get("/projects");

    console.log("Data: ", data);

    projects = data;
  };

  const submitForm = async () => {
    error = "";
    success = "";
    loading = true;

    const payload = {
      name,
      description,
      github_link: githubLink,
      image,
      blog_link: blogLink,
      type,
    };

    if (editMode) {
      payload.id = id;
      try {
        await api.put("/projects", payload);
        name = "";
        description = "";
        githubLink = "";
        image = "";
        blogLink = "";
        type = "";

        success = "Project has been updated successfully";
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
      await api.post("/projects", payload);

      name = "";
      description = "";
      githubLink = "";
      image = "";
      blogLink = "";
      type = "";
      error = "";
      success = "Project uploaded successfully!";
    } catch (err) {
      console.error(err);
      error = err || "Failed to create work experience";
    } finally {
      loading = false;
    }
  };

  onMount(() => {
    getProjects();
  });
</script>

<div class="flex flex-row">
  <Form {submitForm} title="Add Project">
    <Input label="Name" bind:value={name} required={true} />
    <BigInput label="Description" bind:value={description} required={true} />
    <Input label="Github Link" bind:value={githubLink} required={true} />
    <Input label="Blog Link" bind:value={blogLink} />
    <ImageInput
      label="Project Photo"
      onchange={(e) => (image = handleImageChange(e.target.files[0]))}
    />
    <Input label="Type" bind:value={type} required={true} />
    <FormButton {loading} />
    <Notif {error} {success} />
  </Form>
  <DataPreview data={projects} {editHook} {editMode} />
</div>
