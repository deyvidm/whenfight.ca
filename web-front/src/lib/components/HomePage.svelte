<script>
    import currentdudes from "$lib/current-dudes.json";
    import TimeTable from "./TimeTable.svelte";
    import { onMount } from "svelte";

    /**
     * @type {string}
     */
    let defaultWho = "who";
    let everybody = "everybody"
    let who = defaultWho;
    let parsed = Array();
    let isLoaded = true;
    let hideFinished = false;

    onMount(() => {
        who = everybody;
        handleSelectWho();
    });

    function toggleHideFinished() {
        hideFinished = !hideFinished;
    }

    async function handleSelectWho() {
        isLoaded = false;
        if (who == defaultWho) {
            isLoaded = true;
            return;
        }

        let dudes = [who];
        if (who == everybody) {
            dudes = currentdudes;
        }

        const response = await fetch("/api/data", {
            method: "POST",
            body: JSON.stringify(dudes),
            headers: {
                "Content-Type": "application/json",
            },
        });
        parsed = await response.json();
        isLoaded = true;
    }
</script>

<div class="rounded-b-box rounded-se-box relative overflow-x-auto">
    <span></span>
    <div class="card w-150 bg-base-100 shadow-xl space-y-4">
        <div class="card-body">
            <div class="label">
                <span class="label-text">Who's grappling?</span>
            </div>
            <select
                bind:value={who}
                on:change={handleSelectWho}
                class="select select-bordered w-full max-w"
            >
                <option disabled selected value="who">Select a competitor</option>
                <option value="everybody">Everybody</option>
                {#each currentdudes as dude}
                    <option value={dude}>{dude}</option>
                {/each}
            </select>
            <button class="btn btn-primary" on:click={handleSelectWho}>Refresh</button>
            <button class="btn btn-outline" on:click={toggleHideFinished}>
                {#if hideFinished}
                    Show Finished Matches
                {:else}
                    Hide Finished Matches
                {/if}
            </button>
        </div>

        {#if who != defaultWho}
            <TimeTable {isLoaded} {parsed} {hideFinished}></TimeTable>
        {/if}
    </div>
</div>
