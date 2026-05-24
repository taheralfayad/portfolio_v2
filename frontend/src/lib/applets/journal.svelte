<script>
  import "@fontsource/literata";
  import { Heading } from "@lucide/svelte";

  let editor = $state();
  let content = $state([]);
  let scrollHeight = $state();
  let contentIndex = $state(0);

  const insertHeader = (level = 1) => {
    const tag = `h${level}`;

    document.execCommand(
      "insertHTML",
      false,
      `<${tag}>New Header</${tag}><p></p>`,
    );

    editor.focus();
  };

  const moveIndex = (action) => {
    if (action === "previous" && contentIndex - 1 >= 0) {
      contentIndex--;
      editor.innerHTML = content[contentIndex];
    } else if (action === "next" && contentIndex + 1 < content.length) {
      contentIndex++;
      editor.innerHTML = content[contentIndex];
    }
  };

  const updateInput = () => {
    if (!editor) return;

    if (editor.scrollHeight > editor.clientHeight) {
      const overflowNode = editor.lastChild;

      overflowNode?.remove();
      content[contentIndex] = editor.innerHTML;

      contentIndex++;

      if (content[contentIndex] == null) {
        content[contentIndex] = "";
      }

      editor.innerHTML = content[contentIndex];

      if (overflowNode) {
        editor.appendChild(overflowNode);
        content[contentIndex] = editor.innerHTML;
      }

      const range = document.createRange();
      const sel = window.getSelection();
      range.selectNodeContents(editor);
      range.collapse(false);
      sel?.removeAllRanges();
      sel?.addRange(range);
    } else {
      content[contentIndex] = editor.innerHTML;
    }

    scrollHeight = editor.scrollHeight;
  };

  $inspect(content);
</script>

<div class="min-h-screen min-w-screen grid grid-cols-[1fr_2fr_1fr]">
  <div class="border"></div>

  <div class="border relative overflow-hidden flex flex-col">
    <div
      bind:this={editor}
      contenteditable="true"
      oninput={updateInput}
      spellcheck="false"
      class="
        notebook
        literata
        tracking-wider
        flex-1 outline-none
        pl-5 pt-1
        text-xl leading-14
        overflow-auto
        bg-[#fffef8]
        border
        border-r-0
        border-l-0
        border-t-0
        max-h-screen
        whitespace-pre-wrap
        break-words
      "
    ></div>

    <div class="flex justify-between py-2 px-2">
      <button
        class="hover:cursor-pointer"
        onclick={() => moveIndex("previous")}
      >
        &lt-
      </button>
      <p>{contentIndex + 1}</p>
      <button class="hover:cursor-pointer" onclick={() => moveIndex("next")}>
        -&gt
      </button>
    </div>
  </div>

  <div class="flex flex-col items-center justify-start border pt-2">
    <button
      class="border rounded-full p-2 bg-[#fffef8]"
      onclick={() => insertHeader(2)}
    >
      <Heading size={48} />
    </button>
  </div>
</div>

<style>
  .literata {
    font-family: "Literata";
  }
  .notebook :global(h2) {
    font-size: 1.5rem;
    font-weight: bold;
    margin: 0.5rem 0;
  }
  .notebook :global(p) {
    margin: 0.5rem 0;
  }
</style>
