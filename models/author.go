package models

type Author struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type AuthorPrimaryKey struct {
	Id string `json:"id"`
}

type CreateAuthor struct {
	Name string `json:"name"`
}

type UpdateAuthor struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type GetListAuthorRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListAuthorResponse struct {
	Count  int       `json:"count"`
	Author []*Author `json:"Author"`
}
