<script>
  import { onMount } from 'svelte';

  import Input from "../design-system/input.svelte";
  import BigInput from "../design-system/big_input.svelte";
  import ImageInput from "../design-system/image_input.svelte";
  import Form from "../design-system/form.svelte";
  import FormButton from "../design-system/form_button.svelte";
  import Notif from "../design-system/notif.svelte";
  import DataPreview from "../components/data_preview.svelte";

  import { api } from "../utils/api.svelte.js";

  let id = $state(0);
  let title = $state("");
  let caption = $state("");
  let image = $state("");
  let images = $state([]);
  let editMode = $state(false);

  let loading = $state(false);
  let error = $state("");
  let success = $state("");

  const editHook = (data) => {
    if (!data) {
      id = 0;
      title = "";
      caption = "";
      image = "";
      editMode = false;
      return
    }

    editMode = true;
    id = data.id
    title = data.title;
    caption = data.caption;
  }

  const handleImageChange = async (file) => {
    if (!file) return;
    
    const reader = new FileReader();
    
    reader.onload = (e) => {
      image = e.target.result;
    };
    
    reader.readAsDataURL(file);
  };

  const getImages = async () => {
    const data = await api.get("/images");

    images = data;
  }

  const submitForm = async () => {
    error = "";
    success = "";
    loading = true;

    const payload = {
      title,
      caption,
      image
    }

    if (editMode) {
      payload.id = id
      
      try {
        await api.put("/images", payload)
        id = 0;
        title = "";
        caption = "";
        image = "";
        editMode = false;

        success = "Image been updated!"
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
      await api.post("/images", payload)
      
      title = "";
      caption = "";
      image = "";
      success = "Image created."
    } catch (err) {
      console.error(err)
      error =
        err ||
        "Failed to create work experience";
    } finally {
      loading = false;
    }

  }

  onMount(() => {
    getImages();
  })
</script>


<div class="flex flex-row">
  <Form submitForm={submitForm} editMode={editMode} editHook={editHook} title="Add Image">
    <Input label="Title" bind:value={title} required={true}/>
    <BigInput label="Caption" bind:value={caption} required={true}/>
    <ImageInput label="Image" onchange={(e) => handleImageChange(e.target.files[0])}/>
    <FormButton loading={loading} />
    <Notif
      error={error}
      success={success}
    />
  </Form>
  <DataPreview data={images} editHook={editHook}/>
</div>
