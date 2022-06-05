<script lang="ts">
	import type { User } from '$lib/types/User';

	export let menuItems = [
		['Home', '/'],
		['Other Page', '/other']
	];

	export let user: User;
</script>

<div class="flex justify-start">
	{#each menuItems as item}
		<a
			href={item[1]}
			class="rounded-lg px-3 py-2 text-slate-700 font-medium hover:bg-slate-100 hover:text-slate-900"
			>{item[0]}</a
		>
	{/each}

	<a
		href="/logout"
		class="rounded-lg px-3 py-2 text-slate-700 font-medium hover:bg-slate-100 hover:text-slate-900"
		>Welcome, {user.username}</a
	>
</div>
