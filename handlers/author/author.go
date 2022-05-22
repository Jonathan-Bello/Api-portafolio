package author

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Jonathan-Bello/Api-portafolio/handlers"
	"github.com/Jonathan-Bello/Api-portafolio/models"
	"github.com/Jonathan-Bello/Api-portafolio/models/response"
	"github.com/Jonathan-Bello/Api-portafolio/storage"
	psqlauthor "github.com/Jonathan-Bello/Api-portafolio/storage/psql_author"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// GetAll handler for getting all authors
func GetAll(c echo.Context) error {
	authors, err := psqlauthor.GetAll()
	if err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			res := response.New(response.MsgError, err.Error(), http.StatusBadRequest, nil)
			return c.JSON(http.StatusBadRequest, res)
		}

		res := response.New(response.MsgError, err.Error(), http.StatusInternalServerError, nil)
		return c.JSON(http.StatusInternalServerError, res)
	}

	res := response.New(response.MsgOK, "OK", http.StatusOK, authors)
	return c.JSON(http.StatusOK, res)
}

// GetByID handler for getting author by id
func GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := response.New(response.MsgError, handlers.ErrIdRequired.Error(), http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	author, err := psqlauthor.GetByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res := response.New(response.MsgError, err.Error(), http.StatusNotFound, nil)
			return c.JSON(http.StatusNotFound, res)
		}

		res := response.New(response.MsgError, err.Error(), http.StatusInternalServerError, nil)
		return c.JSON(http.StatusInternalServerError, res)
	}

	res := response.New(response.MsgOK, "OK", http.StatusOK, author)
	return c.JSON(http.StatusOK, res)
}

// Create handler for creating author
func Create(c echo.Context) error {
	var author models.Author
	if err := c.Bind(&author); err != nil {
		res := response.New(response.MsgError, handlers.ErrInvalidBody.Error(), http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	//TODO: Validate author data with zero value

	if err := psqlauthor.Create(author); err != nil {
		res := response.New(response.MsgError, err.Error(), http.StatusInternalServerError, nil)
		return c.JSON(http.StatusInternalServerError, res)
	}

	res := response.New(response.MsgOK, "OK", http.StatusCreated, author)
	return c.JSON(http.StatusCreated, res)
}

// Update handler for updating author
func Update(c echo.Context) error {
	var author models.Author
	if err := c.Bind(&author); err != nil {
		res := response.New(response.MsgError, handlers.ErrInvalidBody.Error(), http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	if author.ID == 0 {
		res := response.New(response.MsgError, handlers.ErrIdRequired.Error(), http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	//TODO: Validate author data with zero value

	if err := psqlauthor.Update(author); err != nil {
		if errors.Is(err, storage.ErrNotFoundUpdate) {
			res := response.New(response.MsgError, err.Error(), http.StatusNotFound, nil)
			return c.JSON(http.StatusNotFound, res)
		}
		res := response.New(response.MsgError, err.Error(), http.StatusInternalServerError, nil)
		return c.JSON(http.StatusInternalServerError, res)
	}

	res := response.New(response.MsgOK, "Author updated", http.StatusOK, nil)
	return c.JSON(http.StatusOK, res)
}

// Delete handler for deleting author
func Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := response.New(response.MsgError, handlers.ErrIdRequired.Error(), http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	if err := psqlauthor.Delete(uint(id)); err != nil {
		if errors.Is(err, storage.ErrNotFoundDelete) {
			res := response.New(response.MsgError, err.Error(), http.StatusNotFound, nil)
			return c.JSON(http.StatusNotFound, res)
		}
		res := response.New(response.MsgError, err.Error(), http.StatusInternalServerError, nil)
		return c.JSON(http.StatusInternalServerError, res)
	}

	res := response.New(response.MsgOK, "Author deleted", http.StatusOK, nil)
	return c.JSON(http.StatusOK, res)
}