import axios from "axios";

export function fetchDudeData(participants) {
    const url = "http://localhost/fetchDudeInfo";
    return axios.post(url, {
            "clubID": "47454",
            "eventID": "13528",
            "participants": participants
        })
}