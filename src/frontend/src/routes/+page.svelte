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
	});
</script>
