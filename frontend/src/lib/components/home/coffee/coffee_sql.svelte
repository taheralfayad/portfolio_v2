<script>
  import * as duckdb from "@duckdb/duckdb-wasm";
  import duckdb_wasm from "@duckdb/duckdb-wasm/dist/duckdb-mvp.wasm?url";
  import mvp_worker from "@duckdb/duckdb-wasm/dist/duckdb-browser-mvp.worker.js?url";
  import duckdb_wasm_eh from "@duckdb/duckdb-wasm/dist/duckdb-eh.wasm?url";
  import eh_worker from "@duckdb/duckdb-wasm/dist/duckdb-browser-eh.worker.js?url";
  import CoffeeTable from "$lib/components/home/coffee/coffee_table.svelte";
  import SqlSchema from "$lib/components/home/coffee/sql_schema.svelte";
  import { onMount, onDestroy, tick } from "svelte";

  const SQL_KEYWORDS = [
    "SELECT",
    "FROM",
    "WHERE",
    "GROUP",
    "ORDER",
    "BY",
    "LIMIT",
    "JOIN",
    "LEFT",
    "RIGHT",
    "INNER",
    "OUTER",
    "VIEW",
    "AS",
    "WITH",
    "ON",
    "ASC",
    "DESC",
  ];

  let { parquetBuffers } = $props();

  let db = $state();
  let conn = $state();
  let ready = $state(false);
  let queryEl = $state("");
  let queryText = $state("");
  let results = $state(null);
  let error = $state(null);
  let coffeeSchema = $state({});
  let roastSchema = $state({});
  let coffeeCupSchema = $state({});

  onMount(async () => {
    const MANUAL_BUNDLES = {
      mvp: {
        mainModule: duckdb_wasm,
        mainWorker: mvp_worker,
      },
      eh: {
        mainModule: duckdb_wasm_eh,
        mainWorker: eh_worker,
      },
    };
    const bundle = await duckdb.selectBundle(MANUAL_BUNDLES);
    const worker = new Worker(bundle.mainWorker);
    const logger = new duckdb.ConsoleLogger();
    db = new duckdb.AsyncDuckDB(logger, worker);
    await db.instantiate(bundle.mainModule, bundle.pthreadWorker);
    conn = await db.connect();

    for (const [filename, buffer] of Object.entries(parquetBuffers)) {
      if (buffer.byteLength === 0) {
        console.error("ArrayBuffer is detached!");
        ready = true;
        return;
      }
      const tableName = filename.replace(".parquet", "");
      await db.registerFileBuffer(filename, buffer);
      await conn.query(
        `CREATE VIEW ${tableName} AS SELECT * FROM parquet_scan('${filename}')`,
      );
    }

    const coffeeSchemaResult = await conn.query(`DESCRIBE coffee;`);
    const roastSchemaResult = await conn.query(`DESCRIBE roast;`);
    const coffeeCupSchemaResult = await conn.query(`DESCRIBE coffee_cup;`);

    coffeeSchema = coffeeSchemaResult.toArray().map((row) => row.toJSON());
    roastSchema = roastSchemaResult.toArray().map((row) => row.toJSON());
    coffeeCupSchema = coffeeCupSchemaResult
      .toArray()
      .map((row) => row.toJSON());

    console.log(coffeeSchema);
    ready = true;
  });

  onDestroy(async () => {
    await conn?.close();
    await db?.terminate();
  });

  const submitQuery = async () => {
    error = null;
    results = null;
    try {
      const result = await conn.query(queryText);
      console.log("result:");
      console.log(result);
      results = result.toArray().map((row) => row.toJSON());
    } catch (e) {
      error = e.message;
    }
  };

  const handleTab = (e) => {
    if (e.key === "Tab") {
      // tab to the next element if the last letter is ;
      const lastLetter = queryText.at(-1);
      if (lastLetter === ";") {
        return;
      }

      e.preventDefault();

      const selection = window.getSelection();
      if (!selection.rangeCount) return;

      const range = selection.getRangeAt(0);
      range.deleteContents();

      const tabNode = document.createTextNode("\t");
      range.insertNode(tabNode);

      range.setStartAfter(tabNode);
      range.setEndAfter(tabNode);
      selection.removeAllRanges();
      selection.addRange(range);
    }
  };

  function getCaretOffset(element) {
    const selection = window.getSelection();

    if (!selection.rangeCount) return 0;

    const range = selection.getRangeAt(0);
    const preCaretRange = range.cloneRange();

    preCaretRange.selectNodeContents(element);
    preCaretRange.setEnd(range.endContainer, range.endOffset);

    return preCaretRange.toString().length;
  }

  function restoreCaret(element, offset) {
    const selection = window.getSelection();
    const range = document.createRange();

    let currentOffset = 0;

    const walker = document.createTreeWalker(
      element,
      NodeFilter.SHOW_TEXT,
      null,
    );

    let node;

    while ((node = walker.nextNode())) {
      const nextOffset = currentOffset + node.textContent.length;

      if (offset <= nextOffset) {
        range.setStart(node, offset - currentOffset);
        range.collapse(true);

        selection.removeAllRanges();
        selection.addRange(range);

        return;
      }

      currentOffset = nextOffset;
    }
  }

  const handleInput = () => {
    const caretOffset = getCaretOffset(queryEl);

    queryText = queryEl.innerText;

    queryEl.innerHTML = highlightSQL(queryText);

    restoreCaret(queryEl, caretOffset);
  };

  const highlightSQL = (text) => {
    let html = text;

    for (const keyword of SQL_KEYWORDS) {
      const regex = new RegExp(`\\b${keyword}\\b`, "gi");

      html = html.replace(
        regex,
        `<span class="text-emerald-800 font-bold">$&</span>`,
      );
    }

    return html;
  };
</script>

<div class="flex flex-col w-full mt-14">
  {#if !ready}
    <p>Loading...</p>
  {:else}
    <div class="flex flex-row gap-12">
      <div class="flex flex-col min-w-56 gap-6">
        <h3 class="text-lg">Schema</h3>
        <SqlSchema table="coffee" schema={coffeeSchema}></SqlSchema>
        <SqlSchema table="roast" schema={roastSchema}></SqlSchema>
        <SqlSchema table="coffee_cup" schema={coffeeCupSchema}></SqlSchema>
      </div>
      <div class="w-full">
        <pre class="w-full border bg-white min-h-64 overflow-auto">
          <code
            class="block py-3 px-4 min-h-64 font-mono text-lg leading-relaxed whitespace-pre break-normal outline-none caret-current [tab-size:4]"
            bind:this={queryEl}
            contenteditable="true"
            spellcheck="false"
            aria-label="SQL query editor"
            aria-multiline="true"
            onkeydown={handleTab}
            oninput={handleInput}
            data-placeholder="SELECT * FROM coffee"></code>
        </pre>
        <div class="flex justify-end mt-2">
          <button
            class="p-4 border-4
              after:content-[''] after:absolute after:inset-0
              after:rounded-xl
              after:translate-x-3 after:translate-y-3
              after:-z-10 bg-skills
              cursor-pointer"
            onclick={submitQuery}
            >Submit
          </button>
        </div>
      </div>
    </div>
  {/if}
  {#if error}
    <p class="text-red-500 mt-2">{error}</p>
  {/if}
</div>

{#if results && results.length > 0}
  <CoffeeTable data={results} />
{:else if results}
  <p class="mt-2 text-gray-500">No results.</p>
{/if}

<style>
  code {
    overflow-wrap: normal;
    tab-size: 4;
    -moz-tab-size: 4;
  }

  code[contenteditable]:empty::before {
    content: attr(data-placeholder);
    color: #9ca3af;
    pointer-events: none;
  }
</style>
