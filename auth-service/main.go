package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/dgrijalva/jwt-go"
    "time"
)

var jwtSecret = []byte("your-secret-key")

func generateToken(w http.ResponseWriter, r *http.Request) {
    claims := jwt.MapClaims{
        "username": "admin",
        "exp": time.Now().Add(time.Hour * 1).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, _ := token.SignedString(jwtSecret)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"token": signedToken})
}

func main() {
    http.HandleFunc("/login", generateToken)
    fmt.Println("Auth Service is running on port 8000...")
    http.ListenAndServe(":8000", nil)
}
