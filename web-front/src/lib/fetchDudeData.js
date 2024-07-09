import mock from "$lib/mock.json";

export async function fetchDudeData(dudes) {
    // return { matches: mock } // mock
    // const url = "http://web-back:8080/fetchDudeInfo";
    const url = "http://localhost:8080/fetchDudeInfo";
        try {
            const response = await fetch(url, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    "clubID": "47454",
                    "eventID": "13528",
                    "participants": dudes,
                })
            });

            if (!response.ok) {
                throw new Error('Network response was not ok');
              }

            const data = await response.json();
            // console.log("fetchDudeData", JSON.stringify(data))
            return JSON.stringify(data)
        } catch (error) {
            return { error: error }
        }
}