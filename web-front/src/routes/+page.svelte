<script>
    import HomePage from "$lib/components/HomePage.svelte";

    let who = "everybody";
	
    /** @type {import('./$types').PageData} */
	export let data;

</script>

<div
    class="rounded-b-box rounded-se-box relative overflow-x-auto flex items-center"
>
    <HomePage parsed={data.data} bind:who={who}></HomePage>
</div>
