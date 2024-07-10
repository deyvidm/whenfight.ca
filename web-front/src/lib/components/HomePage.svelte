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

    function toggleHideFinished() {
        hideFinished = !hideFinished;
    }

    function refresh() {
        parsed = parsed.sort((a, b) => {
            return (
                new Date(a.isodate).getTime() - new Date(b.isodate).getTime()
            );
        });
    }

    async function handleSelectWho() {
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

        const response = await fetch("/api/data", {
            method: "POST",
            body: JSON.stringify(dudes),
        });

        const matches = await response.json();

        parsed = [...parsed, ...matches];
        refresh();
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
                <option disabled selected value="who">Who's grappling?</option>
                <option value="everybody">Everybody</option>
                {#each currentdudes as dude}
                    <option value={dude}>{dude}</option>
                {/each}
            </select>
            <button class="btn btn-primary">Refresh</button>
            <button class="btn btn-outline" on:click={toggleHideFinished}>
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
