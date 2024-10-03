
import (
    "net/http"
    "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/schedule", GetSchedule).Methods("GET")
    router.HandleFunc("/tips", GetTips).Methods("GET")
    router.HandleFunc("/optimize", OptimizeRoute).Methods("POST")
    return router
}

func GetSchedule(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Collection schedule"))
}

func GetTips(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Recycling tips"))
}

func OptimizeRoute(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Optimized route"))