<script>
  import { goto } from "$app/navigation";
  import Hero from "$lib/components/home/hero.svelte";
  import WorkProjects from "$lib/components/home/home/work_projects.svelte";
  import PersonalProjects from "$lib/components/home/home/personal_projects.svelte";
  import ExperienceFlow from "$lib/components/home/home/experience_flow.svelte";
  import SkillsTable from "$lib/components/home/home/skills_table.svelte";
  import NavBar from "$lib/components/home/navbar.svelte";
  import DropdownTextfield from "$lib/components/home/coffee/dropdown_textfield.svelte";
  import Gauge from "$lib/components/home/coffee/gauge.svelte";
  import CoffeeDetails from "$lib/components/home/coffee/coffee_details.svelte";
  import CoffeeTable from "$lib/components/home/coffee/coffee_table.svelte";
  import { onMount } from "svelte";

  import { api } from "$lib/utils/api.svelte.js";
  import { formatDate } from "$lib/utils/utils.svelte";

  import coffeesNotFound from "$lib/assets/coffees_not_found.png";

  let workExperiences = $state([]);
  let workProjects = $state([]);
  let personalProjects = $state([]);
  let skills = $state([]);
  let images = $state([]);
  let coffees = $state([]);
  let coffeeCups = $state([]);
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
    const coffeesData = await api.get("/coffees");

    coffees = coffeesData.map((datum) => {
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

    selectedCoffee = coffees[0];

    const coffeeCupData = await api.get(
      `/coffee-cups?coffee_id=${selectedCoffee.id}`,
    );

    coffeeCups = cleanCoffeeCupData(coffeeCupData);
  };

  const selectSuggestion = async (suggestion) => {
    selectedCoffee = coffees.find((item) => item.name === suggestion);
    isFocused = false;

    const coffeeCupData = await api.get(
      `/coffee-cups?coffee_id=${selectedCoffee.id}`,
    );

    if (coffeeCupData != null && coffeeCupData.length > 0) {
      coffeeCups = cleanCoffeeCupData(coffeeCupData);
    } else {
      coffeeCups = {};
    }
  };

  const cleanCoffeeCupData = (coffeeCupData) => {
    return coffeeCupData.map((datum) => {
      return {
        "Date Drank": formatDate(datum.date_drank),
        Temperature: datum.temperature,
        "Days After Roast": datum.days_after_roast,
        Acidity: datum.acidity,
        Body: datum.body,
        Sweetness: datum.sweetness,
        "Water Type": datum.water_type,
        "Grind Size": datum.grind_size,
        Method: datum.method,
        Rating: datum.rating,
      };
    });
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
      {#if coffees && coffees.length > 0}
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
            class="flex flex-col md:flex-row items-center justify-center gap-12 w-full max-w-6xl mt-4"
          >
            <CoffeeDetails coffee={selectedCoffee} />
            <img
              src={selectedCoffee.imageLink}
              class="w-64 h-64 object-cover rounded-lg shadow-lg"
              alt={selectedCoffee.name}
            />
            <Gauge level={selectedCoffee.roastLevel} />
          </div>
        </Hero>
        <CoffeeTable data={coffeeCups} />
      {:else}
        <img src={coffeesNotFound} />
        <p>Sorry, I haven't added any coffees yet.</p>
      {/if}
    </section>
  {/if}
</div>
