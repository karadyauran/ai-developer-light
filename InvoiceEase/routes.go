package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"./handlers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/invoices", handlers.CreateInvoice).Methods("POST")
	router.HandleFunc("/invoices", handlers.GetInvoices).Methods("GET")
	return router
}