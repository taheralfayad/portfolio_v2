<script>
	import { onMount } from "svelte";
	import Hero from "$lib/components/home/hero.svelte";
	import Carousel from "$lib/design-system/carousel.svelte";
	import ExperienceFlow from "$lib/components/home/home/experience_flow.svelte";
	import Projects from "$lib/components/home/home/projects.svelte";
	import SkillsTable from "$lib/components/home/home/skills_table.svelte";
	import Content from "$lib/content/home.json";

	let { currNavValue } = $props();

	import { api } from "$lib/utils/api.svelte.js";

	let workExperiences = $state([]);
	let workProjects = $state([]);
	let personalProjects = $state([]);
	let skills = $state([]);
	let images = $state([]);

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

	onMount(() => {
		getSkills();
		getWorkExperiences();
		getWorkProjects();
		getPersonalProjects();
		getImages();
	});
</script>

<section class="flex items-center justify-center">
	<Hero header={Content.hero.header} subtitle={Content.hero.subtitle}>
		<Carousel {images} />
	</Hero>
</section>
<h2 class="flex justify-center text-xl text-center">
	{Content.sections.workExperiences}
</h2>
<ExperienceFlow items={workExperiences} />
<h2 class="flex justify-center text-xl text-center">
	{Content.sections.workProjects}
</h2>
<Projects projects={workProjects} />
<h2 class="flex justify-center text-xl mt-4 text-center">
	{Content.sections.personalProjects}
</h2>
<Projects projects={personalProjects} />
<h2 class="flex justify-center text-xl mt-4 text-center">
	{Content.sections.skills}
</h2>
<SkillsTable {skills} />
