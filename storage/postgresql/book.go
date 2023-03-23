package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"app/models"
)

type bookRepo struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) *bookRepo {
	return &bookRepo{
		db: db,
	}
}

func (r *bookRepo) CreateBook(req *models.CreateBook) (string, error) {

	var (
		query string
		id    = uuid.New()
	)

	query = `
		INSERT INTO book(
			id, 
			name, 
			price, 
			author_id
		)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(query,
		id.String(),
		req.Name,
		req.Price,
		req.AuthorId,
	)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *bookRepo) GetByIDBook(req *models.BookPrimaryKey) (*models.Book, error) {

	var (
		query  string
		query1 string
		book   models.Book
	)

	query = `
		SELECT
			bo.id,
			bo.name,
			bo.price,
			bo.author_id
			
		FROM book bo
		WHERE id = $1
	`

	err := r.db.QueryRow(query, req.Id).Scan(
		&book.Id,
		&book.Name,
		&book.Price,
		&book.Author.Id,
	)
	if err != nil {
		return nil, err
	}
	query1 = `
		SELECT
			name
			
		FROM author
		WHERE id = $1
	`

	err = r.db.QueryRow(query1, book.Author.Id).Scan(
		&book.Author.Name,
	)

	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *bookRepo) GetListBook(req *models.GetListBookRequest) (resp *models.GetListBookResponse, err error) {

	resp = &models.GetListBookResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			bo.id,
			bo.name,
			bo.price,
			au.id,
			au.name
		FROM book as bo
		JOIN author as au on au.id = bo.author_id
		
	`

	if len(req.Search) > 0 {
		filter += " AND name ILIKE '%' || '" + req.Search + "' || '%' "
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + offset + limit

	fmt.Println(":::Query:", query)

	rows, err := r.db.Query(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {

		var book models.Book
		err = rows.Scan(
			&resp.Count,
			&book.Id,
			&book.Name,
			&book.Price,
			&book.Author.Id,
			&book.Author.Name,
		)

		if err != nil {
			return nil, err
		}

		resp.Books = append(resp.Books, &book)
	}

	return resp, nil
}

func (r *bookRepo) UpdateBook(req *models.UpdateBook) (*models.Book, error) {
	var (
		query string
	)

	query = `
		UPDATE book set name=$1,price=$2  WHERE id = $3
	`

	_, err := r.db.Exec(query, req.Name,
		req.Price,
		req.Id,
	)

	if err != nil {
		return nil, err
	}
	book, err := r.GetByIDBook(&models.BookPrimaryKey{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (r *bookRepo) DeleteBook(req *models.BookPrimaryKey) (string, error) {

	var query string

	query = `
		DELETE FROM book 
		WHERE id = $1
	`

	_, err := r.db.Exec(query, req.Id)

	if err != nil {
		return "", err
	}

	return "deleted", nil
}
