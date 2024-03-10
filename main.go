package main

import (
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"go-slqlite/db"
	"go-slqlite/handlers"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println("DB Connection is funcekd up")
		panic(err)
	}
}

func main() {
	err := db.InitDb()
	if err != nil {
		fmt.Println("DB Connection is funcekd up")
		panic(err)
	}

	fmt.Println("Server xD")
	mux := http.NewServeMux()
	mux.HandleFunc("/books", handlers.GetAll)
	mux.HandleFunc("/book/byid/{id}", handlers.GetById)
	mux.HandleFunc("/book/byisbn/{isbn}", handlers.GetByIsbn)
	mux.HandleFunc("POST /book", handlers.New)
	mux.HandleFunc("PATCH /book/{id}", handlers.Update)
	mux.HandleFunc("DELETE /book/{id}", handlers.Delete)

	http.ListenAndServe(":3000", mux)

	defer db.DB.Close()
}
