<script lang="ts">
	import { Loader } from '@svelteuidev/core';
	import { onMount } from 'svelte';

	onMount(async () => {
		const response = await fetch(`/api/logout`, {
			method: 'GET'
		});

		if (response.ok) {
			document.location = '/login';
		} else {
			const resolvedResponse = await response?.json();
			error = resolvedResponse.error;
			console.error(error);
		}
	});

	$: error = '';
</script>

Logging out...
<Loader />

{#if error}
	<div class="text-red-600">{error}</div>
{/if}
