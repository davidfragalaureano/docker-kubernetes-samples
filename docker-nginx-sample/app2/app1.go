package main
import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"app2": "pong"}`))
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/v1/ping", get).Methods(http.MethodGet)
    log.Fatal(http.ListenAndServe(":80", r))
}