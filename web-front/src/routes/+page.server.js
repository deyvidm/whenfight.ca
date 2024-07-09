/** @type {import('./$types').Actions} */
export const actions = {
    default: async ({ request }) => {
        const data = await request.formData();
        const participants = data.get('participants');
        console.log("form submit      ", participants);
        return await fetchDudeData(participants);
    }
}


import currentdudes from "$lib/current-dudes.json";
import { fetchDudeData } from "$lib/fetchDudeData.js";

export async function load() {
    return await fetchDudeData(currentdudes);
}