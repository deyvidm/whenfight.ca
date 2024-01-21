<script>
    import HomePage from "$lib/components/HomePage.svelte";
    import { onMount } from "svelte";

    let who = "everybody";
</script>

<div
    class="rounded-b-box rounded-se-box relative overflow-x-auto flex items-center"
>
    <HomePage bind:who={who}></HomePage>
</div>
