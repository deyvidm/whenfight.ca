export async function POST({ request }) {

    const url = "http://web-back:8080/fetchDudeInfo";
    try {
        const body = await request.json();
        const {participants} = body;
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "clubID": "47454",
                "eventID": "13528",
                "participants": participants,
            })
        });
        if (response.ok) {
            const data = await response.json();
            return new Response(JSON.stringify(data), {
                status: 200
            });
        } else {
            console.error("Error:", response.status, response.statusText);
        }
    } catch (error) {
        console.error(error);
        return new Response(JSON.stringify({ error: error}), {
            status: 500
        });
    }
}
