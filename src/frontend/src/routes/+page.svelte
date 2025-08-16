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
</div>
