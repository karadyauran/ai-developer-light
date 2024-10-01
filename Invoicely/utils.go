package utils

import (
	"errors"
	"fmt"
	"./models"
)

func SendEmail(invoice models.Invoice) error {
	if invoice.Status == "sent" {
		return errors.New("invoice already sent")
	}
	fmt.Printf("Sending invoice to client: %s, Amount: %.2f\n", invoice.Client, invoice.Amount)
	invoice.Status = "sent"
	return nil
}

func SendReminderEmail(invoice models.Invoice) error {
	if invoice.Status == "paid" {
		return errors.New("invoice already paid")
	}
	fmt.Printf("Sending reminder for invoice ID: %s to client: %s\n", invoice.ID, invoice.Client)
	return nil
}