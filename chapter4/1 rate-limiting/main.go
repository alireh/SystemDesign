package main

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/time/rate"
)

type Message struct {
	Response    string `json:"response"`
	Description string `json:"description"`
}

func endpointExample(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	writer.WriteHeader(http.StatusOK)

	message := Message{
		Response:    "Successful",
		Description: "You've successfully hit the API endpoint",
	}
	err := json.NewEncoder(writer).Encode(&message)
	if err != nil {
		return
	}
}

func rateLimiterMiddleware(next func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	limiter := rate.NewLimiter(3, 6) // max of 6 requests and then three more requests per second
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !limiter.Allow() {
			writer.Write([]byte("rate limit exceeded "))
			return
		} else {
			endpointExample(writer, request)
		}
	})
}

func main() {
	http.HandleFunc("/home", rateLimiterMiddleware(endpointExample))
	err := http.ListenAndServe(":8041", nil)
	if err != nil {
		log.Println("There was an error listening on port :8041", err)
	}
}

//https://www.makeuseof.com/rate-limiting-go-applications/
