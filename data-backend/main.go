package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type Course struct {
	Name   string   `json:"course_name"`
	Grades []string `json:"grades"`
}

type UserGrades struct {
	Username string   `json:"username"`
	Courses  []Course `json:"courses"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type UserInfo struct {
	Username        string   `json:"username"`
	DocumentID      string   `json:"document_id"`
	PendingPayments []string `json:"pending_payments"`
}

// Handler function to return JSON response
func gradesHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the JWT from the Authorization header
	tokenStr := r.Header.Get("Authorization")[len("Bearer "):]
	log.Printf("path: %s token: %s", r.URL.Path, tokenStr)
	claims := &Claims{}

	// Parse the JWT token (we dont have secret, Bob has to think how to distribute it between containers :))
	_, _, err := new(jwt.Parser).ParseUnverified(tokenStr, claims)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	username := claims.Username
	log.Printf("username: %s", username)

	response := fakeRepo(username)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the JWT from the Authorization header
	tokenStr := r.Header.Get("Authorization")[len("Bearer "):]
	log.Printf("path: %s token: %s", r.URL.Path, tokenStr)
	claims := &Claims{}

	// Parse the JWT token (we dont have secret, Bob has to think how to distribute it between containers :))
	_, _, err := new(jwt.Parser).ParseUnverified(tokenStr, claims)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	username := claims.Username
	log.Printf("username: %s", username)

	response := fakeInfoRepo(username)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/grades", gradesHandler)
	http.HandleFunc("/info", infoHandler)
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fakeRepo(username string) UserGrades {
	if username == "userA" {
		return UserGrades{
			Username: username,
			Courses: []Course{
				{Name: "BOT", Grades: []string{"A", "B", "C"}},
				{Name: "Advanced Pentesting", Grades: []string{"B", "A", "A"}},
			},
		}
	} else if username == "userB" {
		return UserGrades{
			Username: username,
			Courses: []Course{
				{Name: "BOT", Grades: []string{"A", "C", "C"}},
				{Name: "Advanced Pentesting", Grades: []string{"A"}},
			},
		}
	} else {
		return UserGrades{
			Username: username,
			Courses: []Course{
				{Name: "BOT", Grades: []string{}},
				{Name: "Advanced Pentesting", Grades: []string{}},
			},
		}
	}
}

func fakeInfoRepo(username string) UserInfo {
	if username == "userA" {
		return UserInfo{
			Username:        username,
			DocumentID:      "A123456789",
			PendingPayments: []string{"Dorm Fee: $200", "Library Fine: $15", "Lab Fee: $50"},
		}
	} else if username == "userB" {
		return UserInfo{
			Username:        username,
			DocumentID:      "B987654321",
			PendingPayments: []string{"Dorm Fee: $250", "Gym Membership: $40"},
		}
	} else {
		return UserInfo{
			Username:        username,
			DocumentID:      "C000000000",
			PendingPayments: []string{"Cafeteria Due: $30"},
		}
	}
}
