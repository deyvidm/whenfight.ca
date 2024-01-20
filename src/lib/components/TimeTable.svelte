<script>
    export let parsed = Array();
    export let isLoaded = false;
</script>

<div class="rounded-b-box rounded-se-box relative overflow-x-auto">
    <span></span>
    <div class="card w-150 bg-base-100 shadow-xl space-y-4">
        <table class="table table-zebra table-sm table-auto">
            {#if parsed.length > 0 && isLoaded}
                <thead>
                    <tr>
                        <th></th>
                        <th>Mat</th>
                        <th>ETA</th>
                        <th>Our Guy</th>
                        <th>Opponent</th>
                    </tr>
                </thead>
                <tbody>
                    {#each parsed as entry}
                        <tr>
                            <td>{entry.style}</td>
                            <td>{entry.day}<br />{entry.mat}</td>
                            <td>{entry.eta}</td>
                            <td>{entry.home_participant}</td>
                            <td>{entry.opponent}</td>
                        </tr>
                    {/each}
                </tbody>
            {:else}
                <div
                class="flex justify-center items-center p-5"
            >
                <div class="loading loading-spinner loading-lg"></div>
            </div>
            {/if}
        </table>
    </div>
</div>
