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
	import { getProjects, type Project } from "$lib/types/project.svelte";
	import { getSkills, type Skill } from "$lib/types/skill.svelte";

	let workExperiences: WorkExperience[] = $state([]);
	let workProjects: Project[] = $state([]);
	let personalProjects: Project[] = $state([]);
	let skills: Skill[] = $state([]);
	let images: Image[] = $state([]);

	const retrieveSkills = async () => {
		const resp = await getSkills();

		if (!resp.success) {
			console.error("something went wrong");
			return;
		}

		skills = resp.resp;
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

	const retrieveWorkProjects = async () => {
		const resp = await getProjects("work", 5);

		if (!resp.success) {
			console.error(resp.error);
			return;
		}

		workProjects = resp.resp;
	};

	const retrievePersonalProjects = async () => {
		const resp = await getProjects("personal", 5);

		if (!resp.success) {
			console.error(resp.error);
			return;
		}

		personalProjects = resp.resp;
	};

	onMount(() => {
		retrieveSkills();
		retrieveWorkExperiences();
		retrieveWorkProjects();
		retrievePersonalProjects();
		retrieveImages();
	});
</script>

{#snippet sectionHeader(text: string)}
	<h2 class="flex mb-6">
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
		{@render sectionHeader(Content.sections.personalProjects)}
		<Projects projects={personalProjects} />
	</div>

	<div>
		{@render sectionHeader(Content.sections.skills)}
		<SkillsTable {skills} />
	</div>
</section>
