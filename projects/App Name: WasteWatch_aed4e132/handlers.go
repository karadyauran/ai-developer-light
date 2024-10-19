
import (
    "encoding/json"
    "net/http"
)

func donateHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        var donation Donation
        if err := json.NewDecoder(r.Body).Decode(&donation); err == nil {
            saveDonation(donation)
            w.WriteHeader(http.StatusCreated)
            return
        }
    }
    w.WriteHeader(http.StatusBadRequest)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        var request FoodRequest
        if err := json.NewDecoder(r.Body).Decode(&request); err == nil {
            saveRequest(request)
            w.WriteHeader(http.StatusCreated)
            return
        }
    }
    w.WriteHeader(http.StatusBadRequest)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        donations := getAllDonations()
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(donations)
        return
    }
    w.WriteHeader(http.StatusMethodNotAllowed)