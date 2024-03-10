package repo

import (
	"go-slqlite/db"
	"go-slqlite/models"
)

func GetOneById(id string) (*models.Book, error) {
	row := db.DB.QueryRow("SELECT id, title, autthor, isbn FROM books where id = ?", id)
	book := models.Book{}
	err := row.Scan(&book.Id, &book.Title, &book.Author, &book.Isbn)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func GetOneByIsbn(isbn string) (*models.Book, error) {
	row := db.DB.QueryRow("SELECT id, title, autthor, isbn FROM books where id = ?", isbn)
	book := models.Book{}
	err := row.Scan(&book.Id, &book.Title, &book.Author, &book.Isbn)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func GetAll() ([]models.Book, error) {
	rows, err := db.DB.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := []models.Book{}
	for rows.Next() {
		b := models.Book{}
		err = rows.Scan(&b.Id, &b.Title, &b.Author, &b.Isbn)
		if err != nil {
			return nil, err
		}
		res = append(res, b)
	}

	return res, nil
}

func New(book models.Book) (*int64, error) {
	insertBook := `INSERT INTO books(title, autthor, isbn ) VALUES (?, ?, ?)`
	statement, err := db.DB.Prepare(insertBook)
	if err != nil {
		return nil, err
	}
	row, err := statement.Exec(book.Title, book.Author, book.Isbn)
	if err != nil {
		return nil, err
	}

	insertedId, err := row.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &insertedId, nil
}

func Update(book models.Book) (*int64, error) {
	updateBook := `UPDATE books SET title = ?, autthor = ?, isbn = ? WHERE id = ?`
	statement, err := db.DB.Prepare(updateBook)
	if err != nil {
		return nil, err
	}
	row, err := statement.Exec(book.Title, book.Author, book.Isbn, book.Id)
	if err != nil {
		return nil, err
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		return nil, err
	}
	return &rowAffected, err
}

func Delete(id string) (*int64, error) {
	updateBook := `DELETE FROM books WHERE id = ?`
	statement, err := db.DB.Prepare(updateBook)
	if err != nil {
		return nil, err
	}
	row, err := statement.Exec(id)
	if err != nil {
		return nil, err
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		return nil, err
	}
	return &rowAffected, err
}
