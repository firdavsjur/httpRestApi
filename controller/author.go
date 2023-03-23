package controller

import (
	"app/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *Controller) AuthorController(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		c.CreateAuthor(w, r)
	case "GET":
		c.GetListAuthor(w, r)
	case "PUT":
		c.UpdateAuthor(w, r)
	case "DELETE":
		c.DeleteAuthor(w, r)
	}

}

func (c *Controller) CreateAuthor(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.HandleFuncResponse(w, "Create Author", 400, err.Error())
		return
	}

	var createAuthor models.CreateAuthor

	err = json.Unmarshal(body, &createAuthor)
	if err != nil {
		c.HandleFuncResponse(w, "Create Author unmarshal json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := c.store.Author().CreateAuthor(&createAuthor)
	if err != nil {
		c.HandleFuncResponse(w, "Storage create Author", http.StatusInternalServerError, err.Error())
		return
	}

	Author, err := c.store.Author().GetByIDAuthor(&models.AuthorPrimaryKey{Id: id})
	if err != nil {
		c.HandleFuncResponse(w, "Storage get by id Author", http.StatusInternalServerError, err.Error())
		return
	}

	body, err = json.Marshal(Author)
	if err != nil {
		c.HandleFuncResponse(w, "Storage get by id marshal Author", http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

func (c *Controller) GetListAuthor(w http.ResponseWriter, r *http.Request) {

	var (
		val    = r.URL.Query()
		limit  int
		offset int
		search string
		err    error
	)

	if _, ok := val["limit"]; ok {
		limit, err = strconv.Atoi(val["limit"][0])
		if err != nil {
			c.HandleFuncResponse(w, "Get List Author limit", http.StatusBadRequest, err.Error())
			return
		}
	}

	if _, ok := val["offset"]; ok {

		offset, err = strconv.Atoi(val["offset"][0])
		if err != nil {
			c.HandleFuncResponse(w, "Get List Author offset", http.StatusBadRequest, err.Error())
			return
		}
	}

	if _, ok := val["search"]; ok {
		search = val["search"][0]
	}

	Authors, err := c.store.Author().GetListAuthor(&models.GetListAuthorRequest{
		Limit:  limit,
		Offset: offset,
		Search: search,
	})
	if err != nil {
		c.HandleFuncResponse(w, "Storage get list Author", http.StatusInternalServerError, err.Error())
		return
	}

	body, err := json.Marshal(Authors)
	if err != nil {
		c.HandleFuncResponse(w, "Storage get list marshal Author", http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (c *Controller) UpdateAuthor(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.HandleFuncResponse(w, "Update Author", 400, err.Error())
		return
	}

	var updateAuthor models.UpdateAuthor

	err = json.Unmarshal(body, &updateAuthor)
	if err != nil {
		c.HandleFuncResponse(w, "Update Author unmarshal json", http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(updateAuthor)
	data, err := c.store.Author().UpdateAuthor(&updateAuthor)
	if err != nil {
		c.HandleFuncResponse(w, "Storage Update Author", http.StatusInternalServerError, err.Error())
		return
	}

	body, err = json.Marshal(data)
	if err != nil {
		c.HandleFuncResponse(w, "Storage get by id marshal Author", http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

func (c *Controller) DeleteAuthor(w http.ResponseWriter, r *http.Request) {

	var id models.AuthorPrimaryKey = models.AuthorPrimaryKey{
		Id: r.URL.Query().Get("id"),
	}

	result, err := c.store.Author().DeleteAuthor(&id)
	if err != nil {
		c.HandleFuncResponse(w, "Storage create Author", http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(result))
}
