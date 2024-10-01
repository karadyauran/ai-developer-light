package handlers

import (
	"encoding/json"
	"net/http"
	"./models"
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	models.AddInvoice(invoice)
	w.WriteHeader(http.StatusCreated)
}

func SendInvoice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req models.SendInvoiceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	invoice, err := models.GetInvoiceByID(req.InvoiceID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = utils.SendEmail(invoice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ListInvoices(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	invoices := models.GetAllInvoices()
	json.NewEncoder(w).Encode(invoices)
}

func SendReminder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req models.SendReminderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	invoice, err := models.GetInvoiceByID(req.InvoiceID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = utils.SendReminderEmail(invoice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}