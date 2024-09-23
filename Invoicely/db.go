package db

import (
	"errors"
	"sync"

	"./models"
)

var (
	invoices = []models.Invoice{}
	mutex    sync.Mutex
	nextID   = 1
)

func SaveInvoice(invoice models.Invoice) error {
	mutex.Lock()
	defer mutex.Unlock()

	invoice.ID = nextID
	nextID++
	invoices = append(invoices, invoice)
	return nil
}

func GetInvoiceByID(id int) (models.Invoice, error) {
	mutex.Lock()
	defer mutex.Unlock()

	for _, invoice := range invoices {
		if invoice.ID == id {
			return invoice, nil
		}
	}
	return models.Invoice{}, errors.New("invoice not found")
}

func GetAllInvoices() []models.Invoice {
	mutex.Lock()
	defer mutex.Unlock()

	return invoices
}

func GenerateReport() models.Report {
	mutex.Lock()
	defer mutex.Unlock()

	var report models.Report
	for _, invoice := range invoices {
		report.TotalIncome += invoice.Amount
		if invoice.Status == "paid" {
			report.Paid++
		} else {
			report.Pending++
		}
	}

	return report
}