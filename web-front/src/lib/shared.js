import axios from "axios";

export function fetchDudeData(participants) {
    console.log("fetching data for", participants);
    const url = "web-back:8080/fetchDudeInfo";
    return axios.post(url, {
            "clubID": "47454",
            "eventID": "13528",
            "participants": participants
        })
}