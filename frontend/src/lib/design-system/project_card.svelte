<script>
	import { Link } from "@lucide/svelte";
	import { Newspaper } from "@lucide/svelte";
	let { project } = $props();
</script>

<div
	class="bg-secondary relative flex flex-col items-center p-8 w-full min-w-[280px] max-w-md"
>
	<div class="w-full h-48 flex items-center justify-center">
		<img
			src={project.image}
			alt={project.name}
			class="max-w-full max-h-full h-auto w-auto object-contain"
		/>
	</div>
	<h3 class="text-xl mt-4 leading-tight">
		{project.name}
	</h3>
	<p class="text-center mt-2 text-sm leading-relaxed">
		{project.description}
	</p>
	<div class="flex w-full justify-center gap-4 mt-auto pt-6">
		<a
			href={project.githubLink}
			class="p-3 hover:cursor-pointer bg-tertiary border"
			target="_blank"
		>
			<Link />
		</a>
	</div>
</div>
