<script>
  import { goto } from "$app/navigation";
  import Hero from "../components/hero.svelte";
  import WorkProjects from "../components/work_projects.svelte";
  import PersonalProjects from "../components/personal_projects.svelte";
  import ExperienceFlow from "../components/experience_flow.svelte";
  import SkillsTable from "../components/skills_table.svelte";
  import NavBar from "../components/navbar.svelte";
  import DropdownTextfield from "../components/dropdown_textfield.svelte";
  import Gauge from "../design-system/gauge.svelte";
  import Carousel from "../design-system/carousel.svelte";
  import CoffeeTable from "../components/coffee_table.svelte";
  import { onMount } from "svelte";

  import { api } from "../utils/api.svelte.js";
  import Navbar from "../components/navbar.svelte";

  let workExperiences = $state([]);
  let workProjects = $state([]);
  let personalProjects = $state([]);
  let skills = $state([]);
  let images = $state([]);
  let coffees = $state([]);
  let selectedCoffee = $state("");
  let searchValue = $state("");
  let isFocused = $state(false);

  let searchBarValues = $derived.by(() => {
    return coffees.map((x) => x.name);
  });

  let suggestions = $derived.by(() => {
    if (!searchValue) return searchBarValues;

    return searchBarValues.filter((item) =>
      item.toLowerCase().includes(searchValue.toLowerCase()),
    );
  });

  let suggestionsHidden = $derived(!isFocused);

  const navItems = [
    {
      value: "Home",
      header: "Hello! My name is Taher Alfayad.",
      subtitle:
        "I'm a husband, older brother, dog dad, Arsenal fan, and software developer.",
      onClick: () => {
        currNavValue = navItems.find((item) => item.value === "Home");
      },
    },
    {
      value: "Coffee",
      header: "Hello! Join me on my coffee journey.",
      subtitle:
        "I love drinking coffee. I've recently started roasting my own, and this is where I document it.",
      onClick: () => {
        currNavValue = navItems.find((item) => item.value === "Coffee");
      },
    },
  ];

  let currNavValue = $derived.by(() => {
    return navItems.find((item) => item.value === "Home");
  });

  const getCurrNavValue = () => {
    return currNavValue.value;
  };

  const getWorkExperiences = async () => {
    const data = await api.get("/work-experiences?limit=3");

    workExperiences = data.map((datum) => ({
      title: datum.title,
      subtitle: datum.workplace,
      description: datum.description,
    }));
  };

  const getWorkProjects = async () => {
    const data = await api.get("/projects?limit=5&type=work");

    workProjects = data.map((datum) => ({
      name: datum.name,
      description: datum.description,
      githubLink: datum.github_link,
      blogLink: datum.blog_link,
      image: datum.image,
      type: datum.type,
    }));
  };

  const getPersonalProjects = async () => {
    const data = await api.get("/projects?limit=5&type=personal");

    personalProjects = data.map((datum) => ({
      name: datum.name,
      description: datum.description,
      githubLink: datum.github_link,
      blogLink: datum.blog_link,
      image: datum.image,
      type: datum.type,
    }));
  };

  const getSkills = async () => {
    const data = await api.get("/skills");

    skills = data.map((datum) => ({
      name: datum.name,
      category: datum.category,
      blogLink: datum.blog_link,
    }));
  };

  const getImages = async () => {
    const data = await api.get("/images");

    images = data.map((datum) => {
      return {
        title: datum.title,
        caption: datum.caption,
        imageLink: datum.image,
      };
    });
  };

  const getCoffees = async () => {
    const data = await api.get("/coffees");

    coffees = data.map((datum) => {
      return {
        id: datum.id,
        name: datum.name,
        roastLevel: datum.roast_level,
        roasterName: datum.roaster_name,
        originCountry: datum.origin_country,
        processing: datum.processing,
        varietal: datum.varietal,
        imageLink: datum.image,
        date: datum.date,
        description: datum.description,
      };
    });
  };

  const selectSuggestion = async (suggestion) => {
    selectedCoffee = suggestion;
    isFocused = false;
    console.log(selectedCoffee);
  };

  onMount(() => {
    getWorkExperiences();
    getWorkProjects();
    getPersonalProjects();
    getSkills();
    getImages();
    getCoffees();
  });

  $inspect(isFocused);
</script>

<Navbar items={navItems} {getCurrNavValue} />
<div class="flex flex-col min-h-screen pt-20">
  {#if currNavValue.value === "Home"}
    <section class="flex items-center justify-center">
      <Hero header={currNavValue.header} subtitle={currNavValue.subtitle}>
        <Carousel {images} />
      </Hero>
    </section>
    <h2 class="flex justify-center text-xl text-center">
      Here's a little something about my work experience:
    </h2>
    <ExperienceFlow items={workExperiences} />
    <h2 class="flex justify-center text-xl text-center">
      Here are some of the projects I've worked on at my job:
    </h2>
    <WorkProjects projects={workProjects} />
    <h2 class="flex justify-center text-xl mt-4 text-center">
      Here are some of the projects I've worked on in my free time:
    </h2>
    <PersonalProjects projects={personalProjects} />
    <h2 class="flex justify-center text-xl mt-4 text-center">
      Here are some technologies that I've used (and loved enough to talk about
      using):
    </h2>
    <SkillsTable {skills} />
  {:else if currNavValue.value === "Coffee"}
    <section class="flex flex-col items-center justify-center">
      <Hero header={currNavValue.header} subtitle={currNavValue.subtitle}>
        <DropdownTextfield
          {suggestionsHidden}
          {suggestions}
          {selectSuggestion}
          currentSuggestion={selectedCoffee}
          bind:searchValue
          onFocus={() => (isFocused = true)}
        />
        <div
          class="flex flex-col md:flex-row items-center justify-center gap-12 w-full max-w-6xl"
        >
          <Carousel {images} />
          <Gauge level={4} />
        </div>
      </Hero>
      <CoffeeTable
        data={[
          { hello: "world", taher: "alfayad" },
          { hello: "hello", alfayad: "taher" },
        ]}
      />
    </section>
  {/if}
</div>
