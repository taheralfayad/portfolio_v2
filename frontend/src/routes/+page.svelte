<script>
  import { goto } from '$app/navigation';
  import Hero from "../components/hero.svelte"; 
  import WorkProjects from "../components/work_projects.svelte";
  import PersonalProjects from "../components/personal_projects.svelte";
  import ExperienceFlow from "../components/experience_flow.svelte";
  import SkillsTable from "../components/skills_table.svelte";
  import NavBar from "../components/navbar.svelte";
  import { onMount } from 'svelte';

  import { api } from "../utils/api.svelte.js";

  let workExperiences = $state([]);
  let workProjects = $state([]);
  let skills = $state([]);

  // let currNavValue = $state("Home");

  // const navBarItems = [
  //   {
  //     "value": "Home",
  //     "onClick": () => {
  //       goto('#');
  //       currNavValue = 'Home';
  //     }
  //   },
  //   {
  //     "value": "Blog",
  //     "onClick": () => {
  //       goto('#')
  //       currNavValue = 'Blog';
  //     }
  //   }
  // ]

  // const getCurrNavValue = () => {
  //   return currNavValue;
  // }

  const getWorkExperiences = async () => {
    const data = await api.get("/work-experiences?limit=3")

    workExperiences = data.map((datum) => ({
      title: datum.title,
      subtitle: datum.workplace,
      description: datum.description,
    }))
  }

  const getWorkProjects = async () => {
    const data = await api.get("/projects?limit=3&type=work")

    workProjects = data.map((datum) => ({
      name: datum.name,
      description: datum.description,
      githubLink: datum.github_link,
      blogLink: datum.blog_link,
      image: datum.image,
      type: datum.type,
    }))
  }

  const getPersonalProjects = async () => {
    const data = await api.get("/projects?limit=3&type=personal")

    personalProjects = data.map((datum) => ({
      name: datum.name,
      description: datum.description,
      githubLink: datum.github_link,
      blogLink: datum.blog_link,
      image: datum.image,
      type: datum.type,
    }))
  }

  const getSkills = async () => {
    const data = await api.get("/skills")

    skills = data.map((datum) => ({
      name: datum.name,
      category: datum.category,
      blogLink: datum.blog_link,
    }))
  }

  onMount(() => {
    getWorkExperiences();
    getWorkProjects();
    getPersonalProjects();
    getSkills();
  });
  
</script>


<div class="flex flex-col">
  <section class="flex items-center justify-center">
    <Hero/>
  </section>
  <h2 class="flex justify-center text-xl text-center">Here's a little something about my work experience:</h2>
  <ExperienceFlow items={workExperiences}/>
  <h2 class="flex justify-center text-xl text-center">Here are some of the projects I've worked on at my job:</h2>
  <WorkProjects projects={workProjects}/>
  <h2 class="flex justify-center text-xl mt-4 text-center">Here are some of the projects I've worked on in my free time:</h2>
  <PersonalProjects/>
  <h2 class="flex justify-center text-xl mt-4 text-center">Here are some technologies that I've used (and loved enough to talk about using):</h2>
  <SkillsTable skills={skills}/>
</div>
