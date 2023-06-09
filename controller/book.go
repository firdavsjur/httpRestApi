package controller

import (
	"app/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *Controller) BookController(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		c.CreateBook(w, r)
	case "GET":
		c.GetListBook(w, r)
	case "PUT":
		c.UpdateBook(w, r)
	case "DELETE":
		c.DeleteBook(w, r)
	}

}

func (c *Controller) CreateBook(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.HandleFuncResponse(w, "Create Book", 400, err.Error())
		return
	}

	var createBook models.CreateBook

	err = json.Unmarshal(body, &createBook)
	if err != nil {
		c.HandleFuncResponse(w, "Create book unmarshal json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := c.store.Book().CreateBook(&createBook)
	if err != nil {
		c.HandleFuncResponse(w, "Storage create book", http.StatusInternalServerError, err.Error())
		return
	}

	book, err := c.store.Book().GetByIDBook(&models.BookPrimaryKey{Id: id})
	if err != nil {
		c.HandleFuncResponse(w, "Storage get by id book", http.StatusInternalServerError, err.Error())
		return
	}

	body, err = json.Marshal(book)
	if err != nil {
		c.HandleFuncResponse(w, "Storage get by id marshal book", http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

func (c *Controller) GetListBook(w http.ResponseWriter, r *http.Request) {

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
			c.HandleFuncResponse(w, "Get List Book limit", http.StatusBadRequest, err.Error())
			return
		}
	}

	if _, ok := val["offset"]; ok {

		offset, err = strconv.Atoi(val["offset"][0])
		if err != nil {
			c.HandleFuncResponse(w, "Get List Book offset", http.StatusBadRequest, err.Error())
			return
		}
	}

	if _, ok := val["search"]; ok {
		search = val["search"][0]
	}

	books, err := c.store.Book().GetListBook(&models.GetListBookRequest{
		Limit:  limit,
		Offset: offset,
		Search: search,
	})
	if err != nil {
		c.HandleFuncResponse(w, "Storage get list book", http.StatusInternalServerError, err.Error())
		return
	}

	body, err := json.Marshal(books)
	if err != nil {
		c.HandleFuncResponse(w, "Storage get list marshal book", http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (c *Controller) UpdateBook(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.HandleFuncResponse(w, "Update Book", 400, err.Error())
		return
	}

	var updateBook models.UpdateBook

	err = json.Unmarshal(body, &updateBook)
	if err != nil {
		c.HandleFuncResponse(w, "Update book unmarshal json", http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(updateBook)
	data, err := c.store.Book().UpdateBook(&updateBook)
	if err != nil {
		c.HandleFuncResponse(w, "Storage Update book", http.StatusInternalServerError, err.Error())
		return
	}

	body, err = json.Marshal(data)
	if err != nil {
		c.HandleFuncResponse(w, "Storage get by id marshal book", http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

func (c *Controller) DeleteBook(w http.ResponseWriter, r *http.Request) {

	var id models.BookPrimaryKey = models.BookPrimaryKey{
		Id: r.URL.Query().Get("id"),
	}

	result, err := c.store.Book().DeleteBook(&id)
	if err != nil {
		c.HandleFuncResponse(w, "Storage create book", http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(result))
}
