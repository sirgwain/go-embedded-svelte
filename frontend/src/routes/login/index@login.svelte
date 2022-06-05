<script lang="ts">
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

<div class="flex items-center justify-center min-h-screen ">
	<div class="px-8 py-6 mt-4 text-left shadow-lg dark:bg-slate-700">
		<h3 class="text-2xl font-bold text-center">Login</h3>

		<form on:submit|preventDefault={onSubmit}>
			<label class="block">
				<span class="text-gray-700 dark:text-slate-400">Username</span>
				<input
					bind:value={username}
					autofocus
					required
					type="text"
					name="username"
					class="mt-1 block rounded-md dark:bg-slate-600 border-gray-300 dark:border-gray-900 shadow-sm dark:focus:border-indigo-300 focus:border-indigo-800 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
				/>
			</label>

			<label class="block">
				<span class="text-gray-700 dark:text-slate-400">Password</span>
				<input
					bind:value={password}
					required
					type="password"
					name="password"
					class="mt-1 block rounded-md dark:bg-slate-600 border-gray-300 dark:border-gray-900 shadow-sm dark:focus:border-indigo-300 focus:border-indigo-800 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
				/>
			</label>
			<button
				class="mt-2 p-2 block rounded-md focus:ring bg-blue-500 border-gray-300 shadow-sm dark:focus:border-indigo-300 focus:border-indigo-800 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
				type="submit">Submit</button
			>
		</form>
		{#if loginError}
			<div class="text-red-600">{loginError}</div>
		{/if}
	</div>
</div>
