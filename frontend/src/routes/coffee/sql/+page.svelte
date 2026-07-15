<script>
	import * as duckdb from "@duckdb/duckdb-wasm";
	import duckdb_wasm from "@duckdb/duckdb-wasm/dist/duckdb-mvp.wasm?url";
	import mvp_worker from "@duckdb/duckdb-wasm/dist/duckdb-browser-mvp.worker.js?url";
	import duckdb_wasm_eh from "@duckdb/duckdb-wasm/dist/duckdb-eh.wasm?url";
	import eh_worker from "@duckdb/duckdb-wasm/dist/duckdb-browser-eh.worker.js?url";
	import CoffeeTable from "$lib/components/home/coffee/coffee_table.svelte";
	import { onMount, onDestroy, tick } from "svelte";
	import LoadingSpinner from "$lib/components/home/coffee/loading_spinner.svelte";
	import { Table } from "@lucide/svelte";

	import { api } from "$lib/utils/api.svelte.js";
	import { ZipReader, BlobReader, BlobWriter } from "@zip.js/zip.js";

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
	let parquetBuffers = $state();

	const getDuckDBFile = async () => {
		const response = await api.download("/duckdbify");

		const zipReader = new ZipReader(new BlobReader(response));
		const entries = await zipReader.getEntries();

		const buffers = {};
		for (const entry of entries) {
			const blob = await entry.getData(new BlobWriter());
			const buffer = await blob.arrayBuffer();
			buffers[entry.filename] = new Uint8Array(buffer);
		}
		await zipReader.close();

		parquetBuffers = buffers;
	};

	onMount(async () => {
		await getDuckDBFile();

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

{#snippet schemaContainer(schemas)}
	{#snippet sqlSchema(table, schema)}
		<details>
			<summary class="list-none flex flex-row cursor-pointer pr-6 gap-4">
				<Table />
				{table}
			</summary>
			<ul>
				{#each schema as scheme}
					<li class="my-2">
						<b>{scheme.column_name}</b> - {scheme.column_type}
					</li>
				{/each}
			</ul>
		</details>
	{/snippet}

	<div
		class="relative w-72 shrink-0 h-144 overflow-y-auto
	  overflow-x-hidden
	  max-w-md
	  bg-tertiary
	"
	>
		<div class="sticky top-4 p-4">
			<h3 class="text-lg mb-4">Schema</h3>

			<div class="flex flex-col gap-6">
				{#each schemas as schema}
					{@render sqlSchema(schema.table, schema.schema)}
				{/each}
			</div>
		</div>
	</div>
{/snippet}

<div class="flex flex-col w-full mt-14 gap-12 h-screen">
	{#if !ready}
		<p>loading...</p>
	{:else}
		<div class="flex flex-row gap-12 h-full">
			{@render schemaContainer([
				{ table: "coffee", schema: coffeeSchema },
				{ table: "roast", schema: roastSchema },
				{ table: "coffee_cup", schema: coffeeCupSchema },
			])}
			<div class="w-full h-full flex flex-col">
				<pre class="w-full bg-white h-64 overflow-auto text-secondary">
          <code
						class="block py-3 px-4 font-mono text-lg leading-relaxed whitespace-pre break-normal outline-none caret-current [tab-size:4]"
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
						class="p-4
							  bg-tertiary
							  cursor-pointer mt-2
							  hover:bg-secondary"
						onclick={submitQuery}
						>Submit
					</button>
				</div>
				{#if results && results.length > 0}
					<CoffeeTable data={results} />
				{:else if results}
					<p class="mt-2">No results.</p>
				{/if}
				{#if error}
					<p class="text-red-500 mt-2">{error}</p>
				{/if}
			</div>
		</div>
	{/if}
</div>

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
