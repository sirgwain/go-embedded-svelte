<script lang="ts">
	import { Button } from '@svelteuidev/core';

	const onSubmit = async () => {
		const data = JSON.stringify({ username, password });

		const response = await fetch(`/api/login`, {
			method: 'post',
			headers: {
				accept: 'application/json'
			},
			body: data
		});

		if (response.ok) {
			document.location = '/';
		} else {
			const resolvedResponse = await response?.json();
			loginError = resolvedResponse.error;
			console.error(loginError);
		}
	};

	let username = '';
	let password = '';
	$: loginError = '';
</script>

<div class="flex items-center justify-center min-h-screen bg-gray-100">
	<div class="px-8 py-6 mt-4 text-left bg-white shadow-lg">
		<h3 class="text-2xl font-bold text-center">Login</h3>

		<form on:submit|preventDefault={onSubmit}>
			<label class="block">
				<span class="text-gray-700">Username</span>
				<input
					bind:value={username}
					autofocus
					required
					type="text"
					name="username"
					class="mt-1 block rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 "
				/>
			</label>

			<label class="block">
				<span class="text-gray-700">Password</span>
				<input
					bind:value={password}
					required
					type="password"
					name="password"
					class="mt-1 block rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 "
				/>
			</label>
			<Button class="mt-2" type="submit">Submit</Button>
		</form>
		{#if loginError}
			<div class="text-red-600">{loginError}</div>
		{/if}
	</div>
</div>
