<script>
    import currentdudes from "$lib/current-dudes.json";
    import TimeTable from "./TimeTable.svelte";
    import { onMount } from "svelte";

    /**
     * @type {string}
     */
    export let who;
    export let parsed = Array();
    let isLoaded = true;
    let hideFinished = false;
    $: if (who) {
        refresh();
    }

    onMount(() => {
        console.log("data", parsed);
        // refresh();
    });

    function refresh() {
        console.log("refresh");
        parsed = parsed.sort((a, b) => {
            return (
                new Date(a.isodate).getTime() - new Date(b.isodate).getTime()
            );
        });
    }

    // async function fetchData() {
    //     isLoaded = false;
    //     if (who == "who") {
    //         isLoaded = true;
    //         return;
    //     }

    //     parsed = [];
    //     let doneDudes = 0;
    //     let dudes = [who];
    //     if (who == "everybody") {
    //         dudes = currentdudes;
    //     }

    //     fetchDudeData(dudes)
    //         .then((resp) => {
    //             parsed = [...parsed, ...resp.data];
    //             parsed = parsed.sort((a, b) => {
    //                 return (
    //                     new Date(a.isodate).getTime() -
    //                     new Date(b.isodate).getTime()
    //                 );
    //             });
    //         })
    //         .finally(() => {
    //             isLoaded = true;
    //         });
    // }

    function toggleValue() {
        hideFinished = !hideFinished;
    }
</script>

<div class="rounded-b-box rounded-se-box relative overflow-x-auto">
    <span></span>
    <div class="card w-150 bg-base-100 shadow-xl space-y-4">
        <div class="card-body">
            <form method="POST">
                <div class="label">
                    <span class="label-text">Who's grappling?</span>
                </div>
                <select
                    name="particpants"
                    bind:value={who}
                    class="select select-bordered w-full max-w"
                >
                    <option disabled selected value="who"
                        >Who's grappling?</option
                    >
                    <option value="everybody">Everybody</option>
                    {#each currentdudes as dude}
                        <option value={dude}>{dude}</option>
                    {/each}
                </select>
                <button class="btn btn-primary">Refresh</button>
            </form>
            <button class="btn btn-outline" on:click={toggleValue}>
                {#if hideFinished}
                    Show Finished Matches
                {:else}
                    Hide Finished Matches
                {/if}
            </button>
        </div>

        <TimeTable {isLoaded} {parsed} {hideFinished}></TimeTable>
    </div>
</div>
