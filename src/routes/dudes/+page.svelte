<script>
    import { onMount } from "svelte";
    import axios from "axios";
    import * as cheerio from "cheerio";
    import "../../app.css";

    let data = "";
    let parsed = Array();
    let matches = Array();
    let page = 1;
    let done = false;
    
    import { goto } from '$app/navigation';

    function goTo() {
        goto('/');
    }

    async function fetchData() {
        do {
            const corsURL = "https://corsproxy.io/?";
            const smoothCompURL =
                "https://grapplingindustries.smoothcomp.com/en/event/13528/schedule/matchlist?search=&club=47454&catid=&mat=&country=&page="+page;
            const url = corsURL + encodeURIComponent(smoothCompURL);
            const response = await axios.get(url);
            const $ = cheerio.load(response.data);

            // select all elems with class 'match-row' inside elems with class 'matches-list'
            const selector = ".matches-list > .match-row";
            matches = $(selector);

            matches.each((i, match) => {
                let participantString = $(match)
                    .find("span.participant")
                    .text();
                let splits = participantString.trim().split(/\n+/);
                let participants = {
                    participant1: splits[0],
                    club1: splits[1],
                    participant2: splits[2],
                    club2: splits[3],
                };
                if (splits[1] == "Current BJJ") {
                    parsed.push(splits[0]);
                } else {
                    parsed.push(splits[2]);
                }
            });
            parsed = [...new Set(parsed)];
            // console.log(parsed);
            page++;
        } while (matches.length > 0);
        parsed = parsed.sort();
        done = true;
    }

    onMount(fetchData);
</script>



<div class="overflow-x-auto">
    <button class="btn" on:click={goTo}>hello</button>


    <table class="table">
        <tbody>
            {#each parsed as entry}
                <tr>
                    <td>{entry}</td>
                </tr>
            {/each}
            {#if done}
                <tr>
                    <td>Done</td>
                </tr>
            {/if}
            {#if !done}
                <tr>
                    <td>Loading...</td>
                </tr>
            {/if}
        </tbody>
    </table>
</div>

<!-- <button class="btn" on:click={fetchData}>Fetch Data</button> -->
