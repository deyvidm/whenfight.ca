package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

type Request struct {
	Participants []string `json:"participants"`
	EventID      string   `json:"eventID"`
	ClubID       string   `json:"clubID"`
}

var redisClient *redis.Client

func main() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Wait for Redis to start
	retryCount := 0
	for {
		err := checkRedisConnection()
		if err == nil {
			break
		}

		if retryCount >= 10 {
			log.Fatal("Failed to connect to Redis after 10 retries, exiting.")
		}

		retryCount++
		log.Println("Waiting for Redis to start...")
		time.Sleep(2 * time.Second)
	}

	http.HandleFunc("/fetchDudeInfo", handleFetchDudeInfo)
	http.ListenAndServe(":8080", nil)
}
func checkRedisConnection() error {
	pong, err := redisClient.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Printf("PONG: %s\n", pong)
	if pong != "PONG" {
		return fmt.Errorf("unexpected response from Redis server: %s", pong)
	}

	return nil
}

func handleFetchDudeInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
		return
	}

	var req Request
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Error unmarshalling JSON",
			http.StatusInternalServerError)
		return
	}

	var results []map[string]interface{}
	var wg sync.WaitGroup

	for _, p := range req.Participants {
		wg.Add(1)
		go func(participant string) {
			defer wg.Done()

			var dudeInfo []map[string]interface{}
			val, err := redisClient.Get(participant).Result()
			if err == redis.Nil {
				fmt.Println("Cache miss for", participant)
				dudeInfo, err = fetchDudeInfo(req.EventID, participant, req.ClubID)
				if err != nil {
					fmt.Printf("Error running Python script for participant %s: %v\n", participant, err)
					return
				}

				val, err := json.Marshal(dudeInfo)
				if err != nil {
					fmt.Printf("Error marshalling JSON for participant %s: %v\n", participant, err)
					return
				}

				err = redisClient.Set(participant, string(val), 10*time.Second).Err()
				if err != nil {
					fmt.Printf("Error setting value in Redis for participant %s: %v\n", participant, err)
					return
				}
			} else if err != nil {
				fmt.Printf("Error getting value from Redis for participant %s: %v\n", participant, err)
				return
			} else {
				fmt.Println("Cache hit for", participant)
				err = json.Unmarshal([]byte(val), &dudeInfo)
				if err != nil {
					fmt.Printf("Error unmarshalling JSON for participant %s: %v\n", participant, err)
					return
				}
			}
			results = append(results, dudeInfo...)
		}(p)
	}

	wg.Wait()

	sort.Slice(results, func(i, j int) bool {
		isodate1 := results[i]["isodate"].(string)
		isodate2 := results[j]["isodate"].(string)
		return isodate1 < isodate2
	})

	jsonResult, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error marshalling result to JSON",
			http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResult)
}
