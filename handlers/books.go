package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go-slqlite/models"
	"go-slqlite/repo"
)

func errorRes(w http.ResponseWriter, err error) {
	errMessage := fmt.Sprintf(`{"Error": "%v"}`, err.Error())
	fmt.Println(errMessage)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(errMessage))
}

func GetById(env *models.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		id := req.PathValue("id")
		book, err := repo.GetOneById(id)
		if err != nil {
			errorRes(w, err)
			return
		}
		json.NewEncoder(w).Encode(book)
	}
}

func GetByIsbn(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	isbn := req.PathValue("isbn")
	book, err := repo.GetOneByIsbn(isbn)
	if err != nil {
		errorRes(w, err)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func GetAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	books, err := repo.GetAll()
	if err != nil {
		errorRes(w, err)
		return
	}
	json.NewEncoder(w).Encode(books)
}

func New(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	book := models.Book{}
	err := json.NewDecoder(req.Body).Decode(&book)
	if err != nil {
		errorRes(w, err)
		return
	}
	id, err := repo.New(book)
	if err != nil {
		errorRes(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	successMessage := fmt.Sprintf(`{"message": "Created id: %v"}`, *id)
	w.Write([]byte(successMessage))
}

func Update(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := req.PathValue("id")
	book := models.Book{}
	err := json.NewDecoder(req.Body).Decode(&book)
	if err != nil {
		errorRes(w, err)
		return
	}
	book.Id, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		errorRes(w, err)
		return
	}
	rowAffected, err := repo.Update(book)
	if err != nil {
		errorRes(w, err)
		return
	}
	successMessage := fmt.Sprintf(`{"message": " Updated %v"}`, rowAffected)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(successMessage))
}

func Delete(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := req.PathValue("id")
	rowAffected, err := repo.Delete(id)
	if err != nil {
		errorRes(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	successMessage := fmt.Sprintf(`{"message": "Deleted %v"}`, rowAffected)
	w.Write([]byte(successMessage))
}
