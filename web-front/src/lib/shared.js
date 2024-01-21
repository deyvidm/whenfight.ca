import axios from "axios";

export function fetchDudeData(participants) {
    console.log("fetching data for", participants);
    const url = "http://34.228.255.157/fetchDudeInfo";
    return axios.post(url, {
            "clubID": "47454",
            "eventID": "13528",
            "participants": participants
        })
}