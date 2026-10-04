<script lang="ts">
	import { Link } from "@lucide/svelte";
	import { type Project } from "$lib/types/project.svelte";

	interface Props {
		projects: Project[];
	}

	let { projects }: Props = $props();
</script>

{#snippet projectCard(project: Project)}
	<div class="flex flex-col min-w-0">
		<div class="w-64 h-48 flex">
			<img
				src={project.image}
				alt={project.name}
				class="max-w-full max-h-full object-contain"
			/>
		</div>
		<a
			class="flex gap-1 underline items-center underline-offset-3"
			href={project.githubLink}
		>
			<h3 class="leading-tight">
				{project.name}
			</h3>
			<Link size={18} />
		</a>
		<p class="text-sm leading-relaxed max-w-4xl">
			{project.description}
		</p>
	</div>
{/snippet}

<section class="flex flex-col items-start justify-start">
	{#each projects as project}
		{@render projectCard(project)}
	{/each}
</section>
