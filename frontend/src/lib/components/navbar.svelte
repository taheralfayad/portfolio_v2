<script>
	import { ChevronDown, ChevronUp, Menu, X } from "@lucide/svelte";
	import { routes } from "$lib/routes/routes.svelte";
	const BUTTON_STYLES = "pr-2 w-10 h-10";
	let openDropdown = $state(null);
	let mobileMenuOpen = $state(false);
	function hasChildren(fn) {
		return typeof fn === "function" && Object.keys(fn).length > 0;
	}
	$effect(() => {
		if (mobileMenuOpen) {
			const scrollY = window.scrollY;
			document.body.style.position = "fixed";
			document.body.style.top = `-${scrollY}px`;
			document.body.style.width = "100%";
			document.body.style.overflow = "hidden";
			return () => {
				document.body.style.position = "";
				document.body.style.top = "";
				document.body.style.width = "";
				document.body.style.overflow = "";
				window.scrollTo(0, scrollY);
			};
		}
	});

	const COLOR_MODE_CLASSES = ["dark", "syria"];
	let currentMode = $state("dark");

	function setColorMode(mode) {
		for (const cls of COLOR_MODE_CLASSES) {
			document.documentElement.classList.toggle(cls, cls === mode);
		}
		currentMode = mode;
	}

	const colorModes = {
		"Color Mode": Object.assign(() => {}, {
			"Dark Mode": () => setColorMode("dark"),
			"Syria Mode": () => setColorMode("syria"),
		}),
	};
</script>

{#snippet navLink(name, fn, kind = "link")}
	{#if kind === "action"}
		<button
			class="flex items-center cursor-pointer w-full px-12 hover:bg-secondary hover:text-font transition h-16"
			onclick={() => {
				fn();
				mobileMenuOpen = false;
			}}
		>
			{name}
		</button>
	{:else}
		<a
			class="flex items-center cursor-pointer w-full px-12 hover:bg-secondary hover:text-font transition h-16"
			href={fn()}
			onclick={() => (mobileMenuOpen = false)}
		>
			{name}
		</a>
	{/if}
{/snippet}

{#snippet navItem(name, fn, kind = "link")}
	<li class="relative flex w-full h-full border-r">
		{#if hasChildren(fn)}
			<div
				class="flex w-full h-full hover:bg-secondary hover:text-font items-center justify-center cursor-pointer"
				onmouseenter={() => (openDropdown = name)}
				onmouseleave={() => (openDropdown = null)}
				role="group"
			>
				{@render navLink(name, fn, kind)}
				{#if openDropdown === name}
					<ul
						class="absolute top-full left-0 flex flex-col bg-tertiary border min-w-full z-10"
					>
						{#each Object.entries(fn) as [childName, childFn]}
							<li class="flex w-full">
								{@render navLink(childName, childFn, kind)}
							</li>
						{/each}
					</ul>
				{/if}
				{#if openDropdown === name}
					<ChevronUp class={BUTTON_STYLES} />
				{:else}
					<ChevronDown class={BUTTON_STYLES} />
				{/if}
			</div>
		{:else}
			{@render navLink(name, fn, kind)}
		{/if}
	</li>
{/snippet}

<nav class="hidden sm:flex justify-between bg-tertiary text-font border-b h-16">
	<ul class="flex h-full items-center">
		{#each Object.entries(routes) as [name, fn]}
			{@render navItem(name, fn, "link")}
		{/each}
	</ul>
	<ul class="flex h-full items-center">
		{#each Object.entries(colorModes) as [name, fn]}
			{@render navItem(name, fn, "action")}
		{/each}
	</ul>
</nav>

<nav
	class="relative flex items-center justify-end bg-tertiary text-font border-b h-16 sm:hidden"
>
	<button
		class="pr-2 cursor-pointer"
		onclick={() => (mobileMenuOpen = !mobileMenuOpen)}
	>
		{#if mobileMenuOpen}
			<X class="w-8 h-8" />
		{:else}
			<Menu class="w-8 h-8" />
		{/if}
	</button>
	{#if mobileMenuOpen}
		<div
			class="fixed inset-x-0 top-16 bottom-0 h-dvh w-full bg-tertiary z-50 overflow-y-auto"
		>
			<ul class="flex flex-col">
				{#each Object.entries(routes) as [name, fn]}
					<li class="flex flex-col w-full border-b">
						{@render navLink(name, fn, "link")}
					</li>
				{/each}
			</ul>
			<div class="flex flex-col w-full px-12 py-4 border-b">
				<label for="mobile-color-mode" class="mb-2"> Color Mode </label>
				<select
					id="mobile-color-mode"
					class="w-full h-12 px-3 bg-tertiary text-font border rounded"
					bind:value={currentMode}
					onchange={(e) => setColorMode(e.currentTarget.value)}
				>
					<option value="dark">Dark Mode</option>
					<option value="syria">Syria Mode</option>
				</select>
			</div>
		</div>
	{/if}
</nav>
