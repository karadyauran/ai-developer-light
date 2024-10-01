package models

import (
	"errors"
	"sync"
)

type Invoice struct {
	ID       string  `json:"id"`
	Client   string  `json:"client"`
	Amount   float64 `json:"amount"`
	DueDate  string  `json:"due_date"`
	Status   string  `json:"status"`
}

type SendInvoiceRequest struct {
	InvoiceID string `json:"invoice_id"`
}

type SendReminderRequest struct {
	InvoiceID string `json:"invoice_id"`
}

var (
	invoices = make(map[string]Invoice)
	mutex    = &sync.Mutex{}
)

func AddInvoice(invoice Invoice) {
	mutex.Lock()
	defer mutex.Unlock()
	invoices[invoice.ID] = invoice
}

func GetInvoiceByID(id string) (Invoice, error) {
	mutex.Lock()
	defer mutex.Unlock()
	invoice, exists := invoices[id]
	if !exists {
		return Invoice{}, errors.New("invoice not found")
	}
	return invoice, nil
}

func GetAllInvoices() []Invoice {
	mutex.Lock()
	defer mutex.Unlock()
	allInvoices := []Invoice{}
	for _, invoice := range invoices {
		allInvoices = append(allInvoices, invoice)
	}
	return allInvoices
}