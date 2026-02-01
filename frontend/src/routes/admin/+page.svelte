<script>
  import WorkExperienceInput from "../../components/work_experience_input.svelte";
  import ProjectsInput from "../../components/projects_input.svelte";
  import NavBar from "../../components/navbar.svelte";
  import { apiRequest } from "../../utils/api.svelte.js";
  import { onMount } from "svelte";

  let currNavValue = $state('');
  let navItems = $state([]);

  const getCurrNavValue = () => {
    return currNavValue;
  }

  const getAllTables = async () => {
    const all_tables = await apiRequest('/all-tables');

    for (const table of all_tables) {
      navItems = [
        ...navItems,
        {
          value: table.table_name,
          onClick: () => {
            currNavValue = table.table_name;
          }
        }
      ];
    }

    currNavValue = all_tables[0].table_name;

  };

  onMount(() => {
    getAllTables();
  });
</script>


<NavBar items={navItems} getCurrNavValue={getCurrNavValue}/>
<div class="flex min-h-screen flex-col pt-20">
  <div class="flex flex-1 items-center justify-center">
    {#if currNavValue === "work_experiences"}
      <WorkExperienceInput/>
    {:else if currNavValue === "projects"}
      <ProjectsInput/>
    {:else}
      <p>not yet implemented</p>
    {/if}
  </div>
</div>
