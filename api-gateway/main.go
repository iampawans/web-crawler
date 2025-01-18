package main

import (
    "fmt"
    "net/http"
    "time"
    "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("your-secret-key")

func validateJWT(w http.ResponseWriter, r *http.Request) {
    tokenString := r.Header.Get("Authorization")
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })

    if err != nil || !token.Valid {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Valid JWT")
}

func main() {
    http.HandleFunc("/validate", validateJWT)
    fmt.Println("API Gateway is running on port 8082...")
    http.ListenAndServe(":8082", nil)
}
