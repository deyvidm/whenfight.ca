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
        // console.log("data", parsed);
        // refresh();
    });

    function refresh() {
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

    async function handleSubmit(event) {
        event.preventDefault();
        
        isLoaded = false;
        if (who == "who") {
            isLoaded = true;
            return;
        }

        parsed = [];
        let dudes = [who];
        if (who == "everybody") {
            dudes = currentdudes;
        }
        
        dudes = ["Michael Canavan"]
        const response = await fetch('/', {
            method: 'POST',
            body: JSON.stringify(dudes),
        })

        const data = await response.json();
        console.log("darta", JSON.parse(JSON.parse(data.data)));


        // parsed = [...parsed, ...data.matches];
        // parsed = parsed.sort((a, b) => {
        //     return (
        //         new Date(a.isodate).getTime() -
        //         new Date(b.isodate).getTime()
        //     );
        // });
        isLoaded = true;
    }

    function toggleValue() {
        hideFinished = !hideFinished;
    }
</script>

<div class="rounded-b-box rounded-se-box relative overflow-x-auto">
    <span></span>
    <div class="card w-150 bg-base-100 shadow-xl space-y-4">
        <div class="card-body">
            <form method="POST" on:submit={handleSubmit}>
                <div class="label">
                    <span class="label-text">Who's grappling?</span>
                </div>
                <select
                    bind:value={who}
                    on:change={() => handleSubmit(event)}
                    class="select select-bordered w-full max-w"
                >
                    <option disabled selected value="who"
                        >Who's grappling?</option>
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

        <di>{parsed}</di>
        <!-- <TimeTable {isLoaded} {parsed} {hideFinished}></TimeTable> -->
    </div>
</div>
