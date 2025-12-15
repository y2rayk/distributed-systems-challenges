package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	log.Printf("Received webhook payload: %s", string(body))
	fmt.Fprintf(w, "Webhook received and logged!")
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)
	port := "8080"
	log.Printf("Starting webhook receiver on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

