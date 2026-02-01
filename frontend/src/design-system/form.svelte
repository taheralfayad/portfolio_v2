<script>
  let { submitForm, title, children } = $props();
</script>

<form
	on:submit|preventDefault={submitForm}
	class="min-w-xl max-w-xl space-y-6 rounded-2xl border border-gray-200 bg-white p-8 shadow-sm"
>
	<h2 class="text-xl font-semibold text-gray-900">
    {title}
	</h2>

  {@render children()}

</form>
