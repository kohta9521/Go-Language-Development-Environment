package mai

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Auther *Author `json:"suthor"`
}

type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	r := mux.NewRouter()

	// エンドポイント
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBooks).Methods("GET")
	r.HandleFunc("/api/books", getBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", getBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", getBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
