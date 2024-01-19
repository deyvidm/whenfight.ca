<script>
    import { onMount } from "svelte";
    import axios from "axios";
    import * as cheerio from 'cheerio';
    import '../app.css';
    import {parseParticipants} from "../lib/shared.js";
    import dudes from '$lib/current-dudes.json'
    
    let parsed = Array();
    let who = "Michael Mitkov";

    async function fetchData() {
        const url = 'https://corsproxy.io/?' + encodeURIComponent('https://grapplingindustries.smoothcomp.com/en/event/13528/schedule/matchlist?search='+who+'&club=47454&catid=0&mat=&country='); 
        const response = await axios.get(url)
        const $ = cheerio.load(response.data);
        // select all elems with class 'match-row' inside elems with class 'matches-list'
        const selector = '.matches-list > .match-row';
        const matches = $(selector);

        parsed = Array();
        matches.each((i, match) => {
            let participantString = $(match).find("span.participant").text()
            const participants = parseParticipants(participantString)

            parsed.push({
                number: $(match).find('.number').text(),
                eta: $(match).find(".eta").text(),
                participants: participants
            })
        });
        parsed = parsed;
        console.log(parsed);
    }
    onMount(fetchData);

</script>

<select bind:value={who} on:change={fetchData} class="select select-bordered w-full max-w-xs">
    <option disabled selected>Who's grappling?</option>
    {#each dudes as dude}
        <option value={dude}>{dude}</option>
    {/each}
</select>

<div class="rounded-b-box rounded-se-box relative overflow-x-auto">
<table class="table">
        <thead>
            <tr>
                <th>Number</th>
                <th>ETA</th>
                <th>Our Guy</th>
                <!-- <th>Club 1</th> -->
                <th>Opponent</th>
                <th>Opponent club</th>
            </tr>
        </thead>
        <tbody>
            {#each parsed as entry}
            <tr>
                <td>{entry.number}</td>
                <td>{entry.eta}</td>
                <td>{entry.participants.p1}</td>
                <!-- <td>{entry.participants.club1}</td> -->
                <td>{entry.participants.p2}</td>
                <td>{entry.participants.club2}</td>
            </tr>
            {/each}
        </tbody>
    </table>
</div>
<!-- <button class="btn" on:click={fetchData}>Fetch Data</button> -->
