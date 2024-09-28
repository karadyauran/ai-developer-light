package inventory

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

var inventory = []Item{
	{ID: 1, Name: "Item1", Stock: 100},
	{ID: 2, Name: "Item2", Stock: 200},
}

func GetInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventory)
}

func AddInventory(w http.ResponseWriter, r *http.Request) {
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID = len(inventory) + 1
	inventory = append(inventory, item)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func UpdateInventory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for index, item := range inventory {
		if item.ID == id {
			inventory = append(inventory[:index], inventory[index+1:]...)
			var updatedItem Item
			_ = json.NewDecoder(r.Body).Decode(&updatedItem)
			updatedItem.ID = id
			inventory = append(inventory, updatedItem)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}
	http.NotFound(w, r)
}

func DeleteInventory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for index, item := range inventory {
		if item.ID == id {
			inventory = append(inventory[:index], inventory[index+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
}