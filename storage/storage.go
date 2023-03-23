package storage

import (
	"app/models"
)

type StorageI interface {
	CloseDB()
	Book() BookRepoI
	Author() AuthorRepoI
}

type BookRepoI interface {
	CreateBook(*models.CreateBook) (string, error)
	GetByIDBook(*models.BookPrimaryKey) (*models.Book, error)
	GetListBook(*models.GetListBookRequest) (*models.GetListBookResponse, error)
	UpdateBook(*models.UpdateBook) (*models.Book, error)
	DeleteBook(*models.BookPrimaryKey) (string, error)
}

type AuthorRepoI interface {
	CreateAuthor(*models.CreateAuthor) (string, error)
	GetByIDAuthor(*models.AuthorPrimaryKey) (*models.Author, error)
	GetListAuthor(*models.GetListAuthorRequest) (*models.GetListAuthorResponse, error)
	UpdateAuthor(*models.UpdateAuthor) (*models.Author, error)
	DeleteAuthor(*models.AuthorPrimaryKey) (string, error)
}
