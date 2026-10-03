<script lang="ts">
	import { type Project } from "$lib/types/project.svelte";

	interface Props {
		projects: Project[];
	}

	let { projects }: Props = $props();
</script>

{#snippet projectCard(project: Project)}
	<div class="flex flex-wrap items-center w-full">
		<div class="flex flex-col flex-1 min-w-0">
			<h3 class="leading-tight">
				{project.name}
			</h3>
			<h4>
				<a class="underline" href={project.githubLink}
					>Link to project</a
				>
			</h4>
			<p class="text-sm leading-relaxed max-w-2xl">
				{project.description}
			</p>
		</div>

		<div class="w-64 h-48 shrink-0 flex items-center justify-center">
			<img
				src={project.image}
				alt={project.name}
				class="max-w-full max-h-full object-contain"
			/>
		</div>
	</div>
{/snippet}

<section class="flex flex-col">
	{#each projects as project}
		{@render projectCard(project)}
	{/each}
</section>
