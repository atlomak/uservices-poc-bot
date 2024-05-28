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

// Handler function to return JSON response
func handler(w http.ResponseWriter, r *http.Request) {
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

func main() {
	http.HandleFunc("/grades", handler)
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
