import urllib.request
from bs4 import BeautifulSoup as soup
from datetime import datetime
import re
import sys

def parseParticipants(splits):
    participants = {
        "home_participant": splits[2],
        "home_club": splits[3],
        "opponent": splits[0],
        "opponent_club": splits[1],
    }
    if splits[1] == "Current BJJ":
        participants = {
            "home_participant": splits[0],
            "home_club": splits[1],
            "opponent": splits[2],
            "opponent_club": splits[3],
        }
    participants["home_participant"] = participants["home_participant"].title()
    participants["opponent"] = participants["opponent"].title()
    return participants

def getCompDates(datestr):
    current_year = datetime.now().year
    datesplit = [s.strip() for s in datestr.text.split("-")]
    compDates = {}
    for d in datesplit:
        date = datetime.strptime(d + " " + str(current_year), "%d %b %Y")
        weekday = date.strftime("%A").lower()
        compDates[weekday] = date
    return compDates

def lambda_handler(event, context):
    headers = { 'User-Agent' : 'Mozilla/5.0 (Windows NT 6.1; Win64; x64)' }
    urls = []
    for participant in event['participants']:
        urls.append("https://grapplingindustries.smoothcomp.com/en/event/{}/schedule/matchlist?search={}&club={}&catid=0&mat=&country=".format(event["eventID"], urllib.parse.quote(participant),  event['clubID']))
    
    parsed = []
    for url in urls:
        request = urllib.request.Request(url, headers=headers)
        response = urllib.request.urlopen(request)
        if response.getcode() > 299:
            print("Error: ", response.getcode(), url)
            continue
        else:
            print("Success: ", response.getcode(), url)

        pageSoup = soup(response.read(), "html.parser")
        compDates = getCompDates(pageSoup.select(".sub-header.date.mute")[0])
        matches = pageSoup.select(".matches-list > .match-row")
        categories = pageSoup.select(".matches-list > .category-row")

        for i, match in enumerate(matches):
            infoSplit = [line for line in match.text.splitlines() if len(line.strip()) > 0][:6]
            participants = parseParticipants(infoSplit[2:6])
            matchInfo = {
                'mat': infoSplit[0],
                'day': re.search(r'(saturday|sunday)', categories[i].text, re.IGNORECASE).group(),
                'style': re.search(r'(no gi|gi)', categories[i].text, re.IGNORECASE).group(),
                'eta': infoSplit[1],
                **participants,  # Merge participants dictionary into the parsed dictionary
            }
            dt = compDates[matchInfo['day'].lower()]
            etaPrecise = datetime.strptime(matchInfo['eta'], "%I:%M %p")
            dt = dt.replace(hour=etaPrecise.hour, minute=etaPrecise.minute)
            matchInfo['isodate'] = dt.isoformat()
            parsed.append(matchInfo)
    
    
    return parsed
    

event = {
    'clubID': '47454',
    'eventID': '13528',
    'participants': sys.argv[1:],
}
res = lambda_handler(event, None)
for r in res:
    print(r)




