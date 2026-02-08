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

  let workExperiences = $state("");

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

  const experience = [
    {
      title: "Applications Programmer I",
      subtitle: "University",
      description: "Django, Canvas LMS, Docker"
    },
    {
      title: "Full Stack Developer",
      subtitle: "EdTech",
      description: "React, Go, LTI 1.3"
    },
    {
      title: "Data & Systems Engineer",
      subtitle: "Analytics",
      description: "AWS Athena, Canvas Data",
      isCurrent: true
    }
  ];

  onMount(() => {
    getWorkExperiences();
  });
  
</script>


<div class="flex flex-col">
  <section class="flex min-h-screen items-center justify-center">
    <Hero/>
  </section>
  <h2 class="flex justify-center text-xl">Here's a little something about my work experience:</h2>
  <ExperienceFlow items={workExperiences}/>
  <h2 class="flex justify-center text-xl">Here are some of the projects I've worked on at my job:</h2>
  <WorkProjects/>
  <h2 class="flex justify-center text-xl mt-4">Here are some of the projects I've worked on in my free time:</h2>
  <PersonalProjects/>
  <h2 class="flex justify-center text-xl mt-4">Here are some technologies that I've used (and loved enough to talk about using):</h2>
  <SkillsTable/>
</div>
