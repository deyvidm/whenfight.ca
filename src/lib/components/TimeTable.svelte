<script>
    import { onMount } from "svelte";
    import axios from "axios";
    import * as cheerio from "cheerio";
    import { parseParticipants } from "$lib/shared.js";
    import dudes from "$lib/current-dudes.json";

    /**
     * @type {string}
     */
     export let who;
    let parsed = Array();
    let isLoading = false;

    $: if (who) {
        fetchData();
    }

    async function fetchData() {
        if (who == "who") {
            return;
        }
        console.log("fetching data for", who)

        isLoading = true;
        const url =
            "https://corsproxy.io/?" +
            encodeURIComponent(
                "https://grapplingindustries.smoothcomp.com/en/event/13528/schedule/matchlist?search=" +
                    who +
                    "&club=47454&catid=0&mat=&country=",
            );
        const response = await axios.get(url);
        const $ = cheerio.load(response.data);
        // select all elems with class 'match-row' inside elems with class 'matches-list'
        const matches = $(".matches-list > .match-row");
        const categories = $(".matches-list > .category-row");

        parsed = Array();
        matches.each((i, match) => {
            let participantString = $(match).find("span.participant").text();
            const participants = parseParticipants(participantString);
            parsed.push({
                mat: $(match).find(".number").text(),
                day: $(categories[i])
                    .text()
                    .match(/saturday|sunday/i),
                dress: $(categories[i])
                    .text()
                    .match(/no gi|gi/i),
                eta: $(match).find(".eta").text(),
                participants: participants,
            });
        });
        parsed = parsed;
        isLoading = false;
    }
    onMount(fetchData);
</script>

<div class="rounded-b-box rounded-se-box relative overflow-x-auto">
    <span></span>
    <div class="card w-150 bg-base-100 shadow-xl space-y-4">
        <div class="artboard phone"></div>

        <select
            bind:value={who}
            class="select select-bordered w-full max-w"
        >
            <option disabled selected value="who">Who's grappling?</option>
            {#each dudes as dude}
                <option value={dude}>{dude}</option>
            {/each}
        </select>
        
        <button class="btn" on:click={fetchData}>Refresh</button>

        <table class="table table-zebra table-sm table-auto">
            <thead>
                <tr>
                    <th></th>
                    <th>Mat</th>
                    <th>ETA</th>
                    <th>Our Guy</th>
                    <th>Opponent</th>
                </tr>
            </thead>

            {#if isLoading}
                <div
                    class="rounded-b-box rounded-se-box relative overflow-x-auto flex justify-center items-center p-5"
                >
                    <span class="loading loading-spinner loading-lg"></span>
                </div>
            {:else}
                <tbody>
                    {#each parsed as entry}
                        <tr>
                            <td>{entry.dress}</td>
                            <td>{entry.day}<br />{entry.mat}</td>
                            <td>{entry.eta}</td>
                            <td>{entry.participants.p1}</td>
                            <td>{entry.participants.p2}</td>
                        </tr>
                    {/each}
                </tbody>
            {/if}
        </table>
    </div>
</div>
