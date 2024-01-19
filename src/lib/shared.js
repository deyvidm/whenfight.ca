/**
 * @param {string} pString
 */
export function parseParticipants(pString) {
    let splits = pString.trim().split(/\n+/);
    let participants = {
        p1 : splits[2],
        club1 : splits[3],
        p2 : splits[0],
        club2 : splits[1],
    }
    if (splits[1] == "Current BJJ") {
        participants = {
            p1 : splits[0],
            club1 : splits[1],
            p2 : splits[2],
            club2 : splits[3],
        }
    }
    return participants
}
