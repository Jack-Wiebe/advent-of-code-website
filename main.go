package main

import (
	"log"
	"net/http"
	"os"

	"advent-of-code-website/handlers"

	"github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

    // API routes
    r.HandleFunc("/", homeHandler).Methods("GET")
    r.HandleFunc("/api/years", handlers.ListYearsHandler).Methods("GET")
    r.HandleFunc("/api/solutions/{year}", handlers.ListSolutionsHandler).Methods("GET")
    r.HandleFunc("/api/run/{year}/{day}", handlers.RunSolutionHandler).Methods("POST")

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Advent of Code Server starting on http://localhost:%s", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./static/index.html")
}