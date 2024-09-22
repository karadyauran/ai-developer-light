package utils

import (
	"errors"
	"sync"
	"./models"
)

var (
	invoices = []models.Invoice{}
	mu       sync.Mutex
)

func SaveInvoice(invoice *models.Invoice) error {
	mu.Lock()
	defer mu.Unlock()

	invoice.ID = generateID()
	invoices = append(invoices, *invoice)
	return nil
}

func FetchInvoices() ([]models.Invoice, error) {
	mu.Lock()
	defer mu.Unlock()

	if len(invoices) == 0 {
		return nil, errors.New("no invoices found")
	}
	return invoices, nil
}

func generateID() string {
	mu.Lock()
	defer mu.Unlock()
	return fmt.Sprintf("INV-%d", len(invoices)+1)
}