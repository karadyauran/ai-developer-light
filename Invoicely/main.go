package main

import (
	"fmt"
	"log"
	"net/http"
	"./handlers"
)

func main() {
	http.HandleFunc("/invoice/create", handlers.CreateInvoice)
	http.HandleFunc("/invoice/send", handlers.SendInvoice)
	http.HandleFunc("/invoice/list", handlers.ListInvoices)
	http.HandleFunc("/invoice/reminder", handlers.SendReminder)

	fmt.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}