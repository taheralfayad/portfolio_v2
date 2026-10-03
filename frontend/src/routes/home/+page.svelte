<script lang="ts">
	import { onMount } from "svelte";
	import Hero from "$lib/components/home/hero.svelte";
	import Carousel from "$lib/design-system/carousel.svelte";
	import WorkExperiences from "$lib/components/home/home/work_experiences.svelte";
	import Projects from "$lib/components/home/home/projects.svelte";
	import SkillsTable from "$lib/components/home/home/skills_table.svelte";
	import Content from "$lib/content/home.json";

	import { api } from "$lib/utils/api.svelte";
	import { getImages, type Image } from "$lib/types/images.svelte";
	import {
		getWorkExperiences,
		type WorkExperience,
	} from "$lib/types/work_experience.svelte";

	let workExperiences: WorkExperience[] = $state([]);
	let workProjects = $state([]);
	let personalProjects = $state([]);
	let skills = $state([]);
	let images: Image[] = $state([]);

	const getSkills = async () => {
		const data = await api.get("/skills");

		skills = data.map((datum) => ({
			name: datum.name,
			category: datum.category,
			blogLink: datum.blog_link,
		}));
	};

	const retrieveImages = async () => {
		const resp = await getImages("home");

		if (resp.success === false) {
			console.error("something went wrong");
			return;
		}

		images = resp.resp;
	};

	const retrieveWorkExperiences = async () => {
		const resp = await getWorkExperiences(3);

		if (!resp.success) {
			console.error(resp.error);
			return;
		}

		workExperiences = resp.resp;
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
		retrieveWorkExperiences();
		getWorkProjects();
		getPersonalProjects();
		retrieveImages();
	});
</script>

{#snippet sectionHeader(text: string)}
	<h2 class="flex text-xl mb-6">
		{text}
	</h2>
{/snippet}

<section class="p-6 flex flex-col gap-12">
	<Hero
		header={Content["home.hero.header"]}
		subtitle={Content["home.hero.subtitle"]}
	>
		<Carousel {images} />
	</Hero>

	<div>
		{@render sectionHeader(Content.sections.workExperiences)}
		<WorkExperiences items={workExperiences} />
	</div>

	<div>
		{@render sectionHeader(Content.sections.workProjects)}
		<Projects projects={workProjects} />
	</div>

	<div>
		<h2 class="flex justify-center text-xl mt-4 text-center">
			{Content.sections.personalProjects}
		</h2>
	</div>
	<Projects projects={personalProjects} />
	<h2 class="flex justify-center text-xl mt-4 text-center">
		{Content.sections.skills}
	</h2>
	<SkillsTable {skills} />
</section>
