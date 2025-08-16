<script>
	import { onMount } from 'svelte';
	const serverUrl = import.meta.env.VITE_SERVER_URL;

	let deviceData;
	let deviceName;
	let deviceTemp;
	let deviceHumidity;

	async function fetchData() {
		try {
			const response = await fetch(serverUrl);
			deviceData = await response.json();

			deviceName = deviceData['deviceName'].toUpperCase();
			deviceTemp = deviceData['temperature'];
			deviceHumidity = deviceData['humidity'];

			console.log(deviceData);
		} catch (err) {
			console.error(err);
		}
	}

	onMount(() => {
		fetchData;

		setInterval(fetchData, 5000);
	});
</script>

<div class="flex h-screen flex-col items-center justify-center gap-10">
	{#if deviceData}
		<h1 class="text-xl">{deviceName}</h1>

		<!-- temperature -->
		<div class="flex flex-row gap-5">
			<div class="flex flex-col items-center justify-center gap-5">
				<div
					class="radial-progress rotate-180"
					style="--value: {deviceTemp}; --size:12rem; --thickness: 1rem;"
					aria-valuenow={deviceTemp}
					role="progressbar"
				></div>
				<h1>Temperature</h1>
				<p>{deviceTemp}°C</p>
			</div>

			<!-- humidity -->
			<div class="flex flex-col items-center justify-center gap-5">
				<div
					class="radial-progress rotate-180"
					style="--value: {deviceHumidity}; --size:12rem; --thickness: 1rem;"
					aria-valuenow={deviceHumidity}
					role="progressbar"
				></div>
				<h1>Humidity</h1>
				<p>{deviceHumidity}%</p>
			</div>
		</div>
	{:else}
		<span class="loading loading-xl loading-spinner"></span>
	{/if}

	<!-- warnings -->
	{#if deviceTemp > 30}
		<div role="alert" class="fixed alert alert-warning bottom-0">
			<svg
				xmlns="http://www.w3.org/2000/svg"
				class="h-6 w-6 shrink-0 stroke-current"
				fill="none"
				viewBox="0 0 24 24"
			>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
				/>
			</svg>
			<span>Warning: Temperature!</span>
		</div>
	{/if}

	{#if deviceHumidity > 60}
		<div role="alert" class="fixed bottom-0 alert alert-warning">
			<svg
				xmlns="http://www.w3.org/2000/svg"
				class="h-6 w-6 shrink-0 stroke-current"
				fill="none"
				viewBox="0 0 24 24"
			>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
				/>
			</svg>
			<span>Warning: Humidity!</span>
		</div>
	{/if}
</div>
