package handlers

import (
	"encoding/json"
	"net/http"

	"./models"
	"./db"
	"./utils"
)

func CreateInvoice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var invoice models.Invoice
	err := json.NewDecoder(r.Body).Decode(&invoice)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = db.SaveInvoice(invoice)
	if err != nil {
		http.Error(w, "Failed to save invoice", http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, invoice)
}

func SendInvoice(w http.ResponseWriter, r *http.Request) {
	// Implementation for sending invoices
}

func TrackInvoice(w http.ResponseWriter, r *http.Request) {
	// Implementation for tracking invoices
}

func GetReports(w http.ResponseWriter, r *http.Request) {
	// Implementation for generating reports
}