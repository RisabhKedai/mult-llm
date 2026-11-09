package main

import (
	"fmt"
	"net/http"

	"mult-llm/backend/backend/chat"
	"mult-llm/backend/backend/health"
)

const PORT = ":8080"

func main() {
	http.HandleFunc("/msg", chat.HandleMessage)
	http.HandleFunc("/health", health.CheckHealth)
	fmt.Printf("Webhook started and listening at %s\n", PORT)
	http.ListenAndServe(PORT, nil)
}
