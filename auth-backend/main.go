package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// UserCredentials holds the structure of user credentials
type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// JWT claims struct
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// jwtKey used to sign tokens
var jwtKey = []byte("my_secret_key")

func main() {
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/user", UserHandler)

	log.Println("Starting app on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// LoginHandler validates user login and returns a JWT
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds UserCredentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// This should be replaced with actual user authentication logic
	if creds.Username != "user" || creds.Password != "pass" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Set expiration time of the token
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the token in the response
	w.Write([]byte(tokenString))
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	tokenStr := r.Header.Get("Authorization")[len("Bearer "):]
	log.Printf("path: %s token: %s", r.URL.Path, tokenStr)
	claims := &Claims{}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Token is valid
	w.Write([]byte(claims.Username))
}
