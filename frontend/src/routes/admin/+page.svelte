<script>
  import WorkExperienceInput from "$lib/components/admin/work_experience_input.svelte";
  import ProjectsInput from "$lib/components/admin/projects_input.svelte";
  import SkillsInput from "$lib/components/admin/skills_input.svelte";
  import UsersInput from "$lib/components/admin/user_input.svelte";
  import ImagesInput from "$lib/components/admin/images_input.svelte";
  import CoffeeInput from "$lib/components/admin/coffee_input.svelte";
  import CoffeeCupInput from "$lib/components/admin/coffee_cup_input.svelte";
  import NavBar from "$lib/components/navbar.svelte";
  import Input from "$lib/design-system/input.svelte";
  import FormButton from "$lib/design-system/form_button.svelte";

  import { api } from "$lib/utils/api.svelte.js";

  let currNavValue = $state("");
  let navItems = $state([]);

  let userName = $state("");
  let password = $state("");
  let accessToken = $state("");
  let loggedIn = $state(false);

  const logIn = async () => {
    const payload = {
      name: userName,
      password: password,
    };

    try {
      const data = await api.post("/login", payload);
      loggedIn = true;
      getAllTables();
    } catch (err) {
      console.error(err);
    }
  };

  const getCurrNavValue = () => {
    return currNavValue;
  };

  const getAllTables = async () => {
    const all_tables = await api.get("/all-tables");

    for (const table of all_tables) {
      navItems = [
        ...navItems,
        {
          value: table.table_name,
          onClick: () => {
            currNavValue = table.table_name;
          },
        },
      ];
    }

    currNavValue = all_tables[0].table_name;
  };
</script>

{#if loggedIn === false}
  <form onsubmit={logIn}>
    <div class="min-h-screen flex items-center justify-center">
      <div>
        <Input label={"Name"} bind:value={userName} />
        <Input label={"Password"} bind:value={password} />
        <div class="pt-4"></div>
        <FormButton loading={false} />
      </div>
    </div>
  </form>
{:else}
  <NavBar items={navItems} {getCurrNavValue} />
  <div class="flex min-h-screen flex-col pt-20">
    <div class="flex flex-1 items-center justify-center">
      {#if currNavValue === "work_experiences"}
        <WorkExperienceInput />
      {:else if currNavValue === "projects"}
        <ProjectsInput />
      {:else if currNavValue === "skills"}
        <SkillsInput />
      {:else if currNavValue === "users"}
        <UsersInput />
      {:else if currNavValue === "images"}
        <ImagesInput />
      {:else if currNavValue === "coffee"}
        <CoffeeInput />
      {:else if currNavValue === "coffee_cup"}
        <CoffeeCupInput />
      {:else}
        <p>not yet implemented</p>
      {/if}
    </div>
  </div>
{/if}
