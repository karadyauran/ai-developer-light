
import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/donate", donateHandler)
    http.HandleFunc("/request", requestHandler)
    http.HandleFunc("/list", listHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))