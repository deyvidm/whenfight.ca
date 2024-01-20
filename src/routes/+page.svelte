<script>
    import HomePage from "$lib/components/HomePage.svelte";
    import { onMount } from "svelte";

    let who = "who"
    $: if (who != "who"){
        console.log("who changed to", who)
        window.sessionStorage.setItem("who", JSON.stringify(who))
    }

    onMount(() => {
        let ses = window.sessionStorage.getItem("who")
        if (ses) {
            console.log("loading sesh", ses)
            who = JSON.parse(ses)
        }
    })
</script>

<div
    class="rounded-b-box rounded-se-box relative overflow-x-auto flex items-center"
>
    <HomePage bind:who={who}></HomePage>
</div>
