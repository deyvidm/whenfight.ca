<script>
    import { onMount } from "svelte";
    import axios from "axios";
    import * as cheerio from 'cheerio';

    let data = "";
    let parsed = Array();

    async function fetchData() {
        const url = 'https://corsproxy.io/?' + encodeURIComponent('https://grapplingindustries.smoothcomp.com/en/event/13528/schedule/matchlist?search=michael+mitkov&club=47454&catid=0&mat=&country='); 
        const response = await axios.get(url)
        const $ = cheerio.load(response.data);
        // select all elems with class 'match-row' inside elems with class 'matches-list'
        const selector = '.matches-list > .match-row';
        const matches = $(selector);

        matches.each((i, match) => {
            let participantString = $(match).find("span.participant").text()
            let splits = participantString.trim().split(/\n+/);
            console.log(splits)
            let participants = {
                "participant1": splits[0],
                "club1": splits[1],
                "participant2": splits[2],
                "club2": splits[3]
            }

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

<button on:click={fetchData}>Fetch Data</button>


<table>
    <thead>
        <tr>
            <th>Number</th>
            <th>ETA</th>
            <th>Participant 1</th>
            <th>Club 1</th>
            <th>Participant 2</th>
            <th>Club 2</th>
        </tr>
    </thead>
    <tbody>
        {#each parsed as entry}
        <tr>
            <td>{entry.number}</td>
            <td>{entry.eta}</td>
            <td>{entry.participants.participant1}</td>
            <td>{entry.participants.club1}</td>
            <td>{entry.participants.participant2}</td>
            <td>{entry.participants.club2}</td>
        </tr>
        {/each}
    </tbody>
</table>


