<script>
  let { error, success} = $props();
</script>

{#if error}
  <p class="text-sm text-red-600">{error}</p>
{/if}

{#if success}
  <p class="text-sm text-green-600">
    {success}
  </p>
{/if}

