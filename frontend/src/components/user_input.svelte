<script>
  import Input from "../design-system/input.svelte";
  import Form from "../design-system/form.svelte";
  import FormButton from "../design-system/form_button.svelte";
  import Notif from "../design-system/notif.svelte";

  import { api } from "../utils/api.svelte.js";

  let name = $state("");
  let password = $state("");
  
  let loading = $state(false)
  let error = $state("")
  let success = $state("")

  const submitForm = async () => {
    error = "";
    success = "";
    loading = true;

    const payload = {
      name,
      password
    }

    try {
      await api.post("/users", payload)

      name = "";
      password = "";
      success = "User created successfully!"
    } catch (err) {
      console.error(err)
      error = err || "Failed to create user"
    } finally {
      loading = false;
    }
  }

</script>

<Form submitForm={submitForm} title={"Add User"}>
  <Input label={"Name"} bind:value={name}/>
  <Input label={"Password"} bind:value={password}/>
  <FormButton loading={loading}/>
  <Notif error={error} success={success}/>
</Form>

