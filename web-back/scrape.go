package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func parseMatchInfo(participant string, infoSplit []string) map[string]string {
	// the last two elems are the participant names; grab em and remove them from splits
	participantNames := []string{infoSplit[len(infoSplit)-1], infoSplit[len(infoSplit)-2]}
	infoSplit = infoSplit[:len(infoSplit)-2]
	pInfo := map[string]string{}

	wonByRegex := regexp.MustCompile(`(?i)(won by)`)
	matRegex := regexp.MustCompile(`\d+-\d+`)
	etaRegex := regexp.MustCompile(`(?i)(\d+:\d+ [ap]m|finished)`)

	c := cases.Title(language.English)
	for i, v := range infoSplit {
		if etaRegex.MatchString(v) {
			pInfo["eta"] = v
		}
		if matRegex.MatchString(v) {
			pInfo["mat"] = v
		}
		if wonByRegex.MatchString(v) {
			pInfo["winner"] = infoSplit[i-2]
			pInfo["result"] = v
		}
		if slices.Contains(participantNames, v) {
			if strings.EqualFold(v, participant) {
				pInfo["home_participant"] = c.String(v)
				pInfo["home_club"] = infoSplit[i+1]
			} else {
				pInfo["opponent"] = c.String(v)
				pInfo["opponent_club"] = infoSplit[i+1]
			}
		}
	}

	return pInfo
}

func getCompDates(datestr string) map[string]time.Time {
	currentYear := time.Now().Year()
	datesplit := strings.Split(datestr, "-")
	compDates := make(map[string]time.Time)
	for _, d := range datesplit {
		date, _ := time.Parse("02 Jan 2006", strings.TrimSpace(d)+" "+strconv.Itoa(currentYear))
		weekday := strings.ToLower(date.Weekday().String())
		compDates[weekday] = date
	}
	return compDates
}

func fetchDudeInfo(eventID, participant, clubID string) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("https://grapplingindustries.smoothcomp.com/en/event/%s/schedule/matchlist?search=%s&club=%s&catid=0&mat=&country=", url.QueryEscape(eventID), url.QueryEscape(participant), url.QueryEscape(clubID))

	fmt.Println("getting", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating GET request: %s\n", err)
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error making GET request to %s: %s\n", url, err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body from %s: %s\n", url, err)
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		fmt.Printf("Error creating document from response body: %s\n", err)
		return nil, err
	}
	parsedMatches := []map[string]interface{}{}
	categories := doc.Find(".matches-list > .category-row")

	compDates := getCompDates(doc.Find(".sub-header.date.mute").Text())
	rawMatches := doc.Find(".matches-list > .match-row")
	fmt.Println("\tfound", rawMatches.Length(), "matches for", participant, "\n\t", url)
	rawMatches.Each(func(i int, s *goquery.Selection) {
		infoSplit := []string{}
		for _, line := range strings.Split(s.Text(), "\n") {
			if len(strings.TrimSpace(line)) > 0 {
				infoSplit = append(infoSplit, line)
			}
		}

		dayRegex := regexp.MustCompile(`(?i)(saturday|sunday)`)
		styleRegex := regexp.MustCompile(`(?i)(gi|no gi)`)
		matchInfo := map[string]interface{}{
			"day":     dayRegex.FindString(categories.Eq(i).Text()),
			"style":   styleRegex.FindString(categories.Eq(i).Text()),
			"isodate": s.Find(".isodate").Text(),
		}

		mInfo := parseMatchInfo(participant, infoSplit)
		for k, v := range mInfo {
			matchInfo[k] = v
		}

		dt := compDates[strings.ToLower(matchInfo["day"].(string))]
		etaPrecise, _ := time.Parse("3:04 pm", matchInfo["eta"].(string))
		dt = dt.Add(time.Hour*time.Duration(etaPrecise.Hour()) + time.Minute*time.Duration(etaPrecise.Minute()))
		matchInfo["isodate"] = dt.Format(time.RFC3339)

		parsedMatches = append(parsedMatches, matchInfo)
	})

	return parsedMatches, nil
}

func mockDudeInfo() ([]map[string]interface{}, error) {
	return []map[string]interface{}{
		{
			"mat":              "4-53",
			"day":              "Sunday",
			"style":            "Gi",
			"eta":              "11:02 am",
			"home_participant": "Michael Mitkov",
			"home_club":        "Current BJJ",
			"opponent":         "Benjamin Keyfitz",
			"opponent_club":    "House of Combat",
			"isodate":          "2024-01-21T11:02:00",
		},
		{
			"mat":              "4-55",
			"day":              "Sunday",
			"style":            "Gi",
			"eta":              "11:09 am",
			"home_participant": "Michael Mitkov",
			"home_club":        "Current BJJ",
			"opponent":         "Benjamin Keyfitz",
			"opponent_club":    "House of Combat",
			"isodate":          "2024-01-21T11:09:00",
		},
		{
			"mat":              "4-57",
			"day":              "Sunday",
			"style":            "Gi",
			"eta":              "11:17 am",
			"home_participant": "Michael Mitkov",
			"home_club":        "Current BJJ",
			"opponent":         "Benjamin Keyfitz",
			"opponent_club":    "House of Combat",
			"isodate":          "2024-01-21T11:17:00",
		},
		{
			"mat":              "7-125",
			"day":              "Sunday",
			"style":            "No Gi",
			"eta":              "3:51 pm",
			"home_participant": "Michael Mitkov",
			"home_club":        "Current BJJ",
			"opponent":         "Kingston Glinsky",
			"opponent_club":    "Warrior Mixed Martial Arts - Newmarket",
			"isodate":          "2024-01-21T15:51:00",
		},
		{
			"mat":              "7-127",
			"day":              "Sunday",
			"style":            "No Gi",
			"eta":              "4:00 pm",
			"home_participant": "Michael Mitkov",
			"home_club":        "Current BJJ",
			"opponent":         "Fionn Craven",
			"opponent_club":    "Xtreme Couture Toronto",
			"isodate":          "2024-01-21T16:00:00",
		},
		{
			"mat":              "7-128",
			"day":              "Sunday",
			"style":            "No Gi",
			"eta":              "4:04 pm",
			"home_participant": "Michael Mitkov",
			"home_club":        "Current BJJ",
			"opponent":         "Damian D'Aloisio",
			"opponent_club":    "Xtreme Couture Toronto",
			"isodate":          "2024-01-21T16:04:00",
		},
	}, nil
}
