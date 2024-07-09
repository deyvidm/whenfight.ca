/** @type {import('./$types').Actions} */
export const actions = {
    default: async ({ request }) => {
        console.log("page server submit action, ")
        const participants = await request.json();
        console.log("participants ", participants);

        
        const resp = await fetchDudeData(participants);
        console.log("---- resp: ----")
        console.log(resp)
        console.log("----")
        return resp

    }
}

import currentdudes from "$lib/current-dudes.json";
import { fetchDudeData } from "$lib/fetchDudeData.js";
import { json } from "@sveltejs/kit";

export async function load() {
    const matches = await fetchDudeData(currentdudes);
    return { props: {matches}}
}