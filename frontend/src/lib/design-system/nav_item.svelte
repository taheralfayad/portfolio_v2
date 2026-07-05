<script>
	let { value, onClick, getCurrNavValue } = $props();
</script>

<li class="flex w-full h-full">
	<button
		class="flex items-center text-sm cursor-pointer w-full h-full px-12 border-r hover:bg-secondary hover:text-font transition"
		onclick={onClick}
	>
		{value}
	</button>
</li>
