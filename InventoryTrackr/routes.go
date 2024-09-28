package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"./inventory"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/inventory", inventory.GetInventory).Methods("GET")
	r.HandleFunc("/inventory", inventory.AddInventory).Methods("POST")
	r.HandleFunc("/inventory/{id}", inventory.UpdateInventory).Methods("PUT")
	r.HandleFunc("/inventory/{id}", inventory.DeleteInventory).Methods("DELETE")
	return r
}