<script lang="ts">
	import { authGuard } from '$lib/authGuard';
	import Menu from '$lib/components/Menu.svelte';
	import type { User } from '$lib/types/User';
	import { onMount } from 'svelte';
	import '../app.css';

	// verify the user, redirect otherwise
	onMount(async () => (user = await authGuard()));

	let user: User | undefined;
</script>

{#if user}
	<Menu {user} />
{/if}

<div class="m-5">
	<slot />
</div>
