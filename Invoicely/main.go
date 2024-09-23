package main

import (
	"log"
	"net/http"

	"./handlers"
)

func main() {
	http.HandleFunc("/create-invoice", handlers.CreateInvoice)
	http.HandleFunc("/send-invoice", handlers.SendInvoice)
	http.HandleFunc("/track-invoice", handlers.TrackInvoice)
	http.HandleFunc("/get-reports", handlers.GetReports)

	log.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}