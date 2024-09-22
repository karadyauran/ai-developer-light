package handlers

import (
	"encoding/json"
	"net/http"
	"./models"
	"./utils"
)

func CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var invoice models.Invoice
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := utils.SaveInvoice(&invoice); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(invoice)
}

func GetInvoices(w http.ResponseWriter, r *http.Request) {
	invoices, err := utils.FetchInvoices()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(invoices)
}