<script>
  import Input from "../design-system/input.svelte";

  let {
    suggestionsHidden,
    suggestions,
    selectSuggestion,
    currentSuggestion,
    searchValue = $bindable(),
    onFocus,
  } = $props();
</script>

<div class="relative w-full max-w-md">
  <Input
    bind:value={searchValue}
    label="Look for a coffee..."
    required={false}
    {onFocus}
  />

  {#if !suggestionsHidden && suggestions.length}
    <ul
      class="
        absolute left-0 right-0 mt-1
        bg-white
        rounded-xl
        border border-gray-200
        shadow-lg
        max-h-56 overflow-y-auto
        z-50
      "
    >
      {#each suggestions as suggestion}
        <li
          class="
            px-4 py-2 text-sm text-gray-700
            cursor-pointer transition
            hover:bg-button/80 hover:font-black
            {currentSuggestion === suggestion ? 'bg-button/80 font-black' : ''}
          "
          role="option"
          tabindex="0"
          onclick={() => selectSuggestion(suggestion)}
        >
          {suggestion}
        </li>
      {/each}
    </ul>
  {/if}
</div>
