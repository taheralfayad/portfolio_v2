<script>
  import ProjectCard from "$lib/design-system/project_card.svelte";

  let { projects } = $props();
</script>

<section class="flex flex-wrap justify-center gap-6 py-8 px-24">
  {#each projects as project}
    <ProjectCard {project} />
  {/each}
</section>
