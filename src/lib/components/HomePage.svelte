<script>
    import { onMount } from "svelte";
    import dudes from "$lib/current-dudes.json";
    import TimeTable from "./TimeTable.svelte";
    import {fetchDudeData} from "$lib/shared.js";

    /**
     * @type {string}
     */
     export let who;
    let parsed = Array();

    $: if (who) {
        refresh()
    }

    function refresh(){
        parsed = [];
        fetchData()
    }

    async function fetchData() {
        if (who == "who") {
            return;
        }
        fetchDudeData([who]).then(resp => {
            parsed = [...resp.data];
        });
    }
    onMount(fetchData);
</script>

<div class="rounded-b-box rounded-se-box relative overflow-x-auto">
    <span></span>
    <div class="card w-150 bg-base-100 shadow-xl space-y-4">
        <select
            bind:value={who}
            class="select select-bordered w-full max-w"
        >
            <option disabled selected value="who">Who's grappling?</option>
            {#each dudes as dude}
                <option value={dude}>{dude}</option>
            {/each}
        </select>
        
        <button class="btn" on:click={refresh}>Refresh</button>

        <TimeTable parsed={parsed} ></TimeTable>
    </div>
</div>
