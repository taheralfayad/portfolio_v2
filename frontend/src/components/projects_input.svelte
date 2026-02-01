<script>
  import Input from "../design-system/input.svelte";
  import BigInput from "../design-system/big_input.svelte";
  import DateInput from "../design-system/date_input.svelte";
  import ImageInput from "../design-system/image_input.svelte";
  import FormButton from "../design-system/form_button.svelte";
  import Notif from "../design-system/notif.svelte";
  import Form from "../design-system/form.svelte";

  import { api } from "../utils/api.svelte.js";

  let name = $state("");
  let description = $state("");
  let githubLink = $state("");
  let image = $state("");
  let blogLink = $state("");
  let type = $state("");

  let error = $state("");
  let success = $state("");
  let loading = $state(false);

  const handleImageChange = async (file) => {
    if (!file) return;
    
    const reader = new FileReader();
    
    reader.onload = (e) => {
      image = e.target.result;
    };
    
    reader.readAsDataURL(file);
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
      type
    }

    try {
      await api.post("/projects", payload)

      name = "";
      description = "";
      githubLink = "";
      image = "";
      blogLink = "";
      type = "";
      error = "";
      success = "Project uploaded successfully!";
      loading = "";

    } catch (err) {
      console.error(err)
      error =
        err ||
        "Failed to create work experience";
    } finally {
      loading = false;
    }
  }

</script>

<Form submitForm={submitForm} title="Add Project">
  <Input label="Name" bind:value={name} />
  <BigInput label="Description" bind:value={description} />
  <Input label="Github Link" bind:value={githubLink} />
  <Input label="Blog Link" bind:value={blogLink} />
  <ImageInput
    label="Project Photo"
    onchange={(e) => handleImageChange(e.target.files[0])}
  />
  <Input label="Type" bind:value={type} />
  <FormButton loading={loading} />
  <Notif
    error={error}
    success={success}
  />
</Form>
