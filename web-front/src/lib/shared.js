import axios from "axios";

export function fetchDudeData(participants) {
    console.log("fetching data for", participants);
    const url = "https://ynh8mrnzla.execute-api.us-east-1.amazonaws.com/fetchDudeInfo";
    return axios.post(url, {
            "clubID": "47454",
            "eventID": "13528",
            "participants": participants
        })
}