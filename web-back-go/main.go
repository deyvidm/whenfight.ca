package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
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
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if err := checkRedisConnection(); err != nil {
		panic(err)
	}
	http.HandleFunc("/fetchDudeInfo", fetchDudeInfo)
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

func fetchDudeInfo(w http.ResponseWriter, r *http.Request) {
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

	val, err := redisClient.Get(req.Participants[0]).Result()
	if err == redis.Nil {
		out, err := exec.Command("python", "scrape.py", req.Participants[0]).Output()
		if err != nil {
			http.Error(w, "Error running Python script",
				http.StatusInternalServerError)
			return
		}

		val = string(out)
		err = redisClient.Set(req.Participants[0], val, 10*time.Second).Err()
		if err != nil {
			http.Error(w, "Error setting value in Redis",
				http.StatusInternalServerError)
			return
		}
	} else if err != nil {
		http.Error(w, "Error getting value from Redis",
			http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, val)
}
