package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

func storeData(w http.ResponseWriter, r *http.Request) {
    var data map[string]interface{}
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    fmt.Printf("Storing data: %v
", data)
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"status": "Data stored!"})
}

func main() {
    http.HandleFunc("/store", storeData)
    fmt.Println("Data Service is running on port 8081...")
    http.ListenAndServe(":8081", nil)
}
