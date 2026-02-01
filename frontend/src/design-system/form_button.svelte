<script>
  let { loading } = $props();
</script>

<button
  type="submit"
  disabled={loading}
  class="inline-flex items-center justify-center rounded-md bg-blue-600 px-4 py-2 text-sm font-semibold text-white shadow-sm transition hover:bg-blue-700 disabled:cursor-not-allowed disabled:opacity-60"
>
  {loading ? "Saving…" : "Save"}
</button>
