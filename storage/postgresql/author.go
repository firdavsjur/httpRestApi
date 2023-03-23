package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"app/models"
)

type authorRepo struct {
	db *sql.DB
}

func NewAuthorRepo(db *sql.DB) *authorRepo {
	return &authorRepo{
		db: db,
	}
}

func (r *authorRepo) CreateAuthor(req *models.CreateAuthor) (string, error) {

	var (
		query string
		id    = uuid.New()
	)

	query = `
		INSERT INTO author(
			id, 
			name
		)
		VALUES ($1, $2)
	`

	_, err := r.db.Exec(query,
		id.String(),
		req.Name,
	)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *authorRepo) GetByIDAuthor(req *models.AuthorPrimaryKey) (*models.Author, error) {

	var (
		query  string
		author models.Author
	)

	query = `
		SELECT
			id,
			name
		FROM author
		WHERE id = $1
	`

	err := r.db.QueryRow(query, req.Id).Scan(
		&author.Id,
		&author.Name,
	)

	if err != nil {
		return nil, err
	}

	return &author, nil
}

func (r *authorRepo) GetListAuthor(req *models.GetListAuthorRequest) (resp *models.GetListAuthorResponse, err error) {

	resp = &models.GetListAuthorResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			name
		FROM author
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

		var author models.Author
		err = rows.Scan(
			&resp.Count,
			&author.Id,
			&author.Name,
		)

		if err != nil {
			return nil, err
		}

		resp.Author = append(resp.Author, &author)
	}

	return resp, nil
}

func (r *authorRepo) UpdateAuthor(req *models.UpdateAuthor) (*models.Author, error) {
	var (
		query string
	)

	query = `
		UPDATE author set name=$1 WHERE id = $2
	`

	_, err := r.db.Exec(query, req.Name,
		req.Id,
	)

	if err != nil {
		return nil, err
	}
	author, err := r.GetByIDAuthor(&models.AuthorPrimaryKey{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return author, nil
}

func (r *authorRepo) DeleteAuthor(req *models.AuthorPrimaryKey) (string, error) {

	var query string

	query = `
		DELETE FROM author 
		WHERE id = $1
	`

	_, err := r.db.Exec(query, req.Id)

	if err != nil {
		return "", err
	}

	return "deleted", nil
}
