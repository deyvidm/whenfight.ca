<script>
    import TimeTable from "$lib/components/TimeTable.svelte";
    import { onMount } from "svelte";
    import dudes from "$lib/current-dudes.json";
    import { fetchDudeData } from "$lib/shared";

    let parsed = [];

    function fetchData(){
        dudes.forEach(dude => {
            fetchDudeData([dude]).then(resp => {
                parsed = [...parsed, ...resp.data];
                parsed = parsed.sort((a,b)=>{
                    return new Date(a.isodate) - new Date(b.isodate);
                })
            });
        }); 
    }

    function refresh(){
        parsed = [];
        fetchData()
    }

    onMount(fetchData);
</script>


<div class="rounded-b-box rounded-se-box relative overflow-x-auto">
    <span></span>
    <div class="card w-150 bg-base-100 shadow-xl space-y-4">       
        <button class="btn" on:click={refresh}>Refresh</button>

        <TimeTable parsed={parsed} ></TimeTable>
    </div>
</div>
