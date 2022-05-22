package blog

import (
	"errors"
	"net/http"

	"github.com/Jonathan-Bello/Api-portafolio/models"
	"github.com/Jonathan-Bello/Api-portafolio/models/response"
	psqlblog "github.com/Jonathan-Bello/Api-portafolio/storage/psql_blog"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// GetAll handler for getting all authors
func GetAll(c echo.Context) error {
	authors, err := psqlblog.GetAll()
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

// Create handler for creating author
func Create(c echo.Context) error {
	var blog models.Blog
	if err := c.Bind(&blog); err != nil {
		res := response.New(response.MsgError, "invalid body to parser json", http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	if err := psqlblog.Create(blog); err != nil {
		res := response.New(response.MsgError, err.Error(), http.StatusInternalServerError, nil)
		return c.JSON(http.StatusInternalServerError, res)
	}

	res := response.New(response.MsgOK, "Blog created", http.StatusCreated, blog)
	return c.JSON(http.StatusCreated, res)
}
