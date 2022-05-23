package blog

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Jonathan-Bello/Api-portafolio/handlers"
	"github.com/Jonathan-Bello/Api-portafolio/models"
	"github.com/Jonathan-Bello/Api-portafolio/models/response"
	"github.com/Jonathan-Bello/Api-portafolio/storage"
	psqlblog "github.com/Jonathan-Bello/Api-portafolio/storage/psql_blog"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// GetAll handler for getting all blogs
func GetAll(c echo.Context) error {
	blogs, err := psqlblog.GetAll()
	if err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			res := response.New(response.MsgError, err.Error(), http.StatusBadRequest, nil)
			return c.JSON(http.StatusBadRequest, res)
		}

		res := response.New(response.MsgError, err.Error(), http.StatusInternalServerError, nil)
		return c.JSON(http.StatusInternalServerError, res)
	}

	res := response.New(response.MsgOK, "OK", http.StatusOK, blogs)
	return c.JSON(http.StatusOK, res)
}

// GetByID handler for getting blog by id
func GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := response.New(response.MsgError, handlers.ErrIdRequired.Error(), http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	blog, err := psqlblog.GetByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res := response.New(response.MsgError, err.Error(), http.StatusNotFound, nil)
			return c.JSON(http.StatusNotFound, res)
		}

		res := response.New(response.MsgError, err.Error(), http.StatusInternalServerError, nil)
		return c.JSON(http.StatusInternalServerError, res)
	}

	res := response.New(response.MsgOK, "OK", http.StatusOK, blog)
	return c.JSON(http.StatusOK, res)
}

// Create handler for creating blog
func Create(c echo.Context) error {
	var blog models.Blog
	if err := c.Bind(&blog); err != nil {
		res := response.New(response.MsgError, "invalid body to parser json", http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	// TODO validate zero values

	if err := psqlblog.Create(blog); err != nil {
		res := response.New(response.MsgError, err.Error(), http.StatusInternalServerError, nil)
		return c.JSON(http.StatusInternalServerError, res)
	}

	res := response.New(response.MsgOK, "Blog created", http.StatusCreated, blog)
	return c.JSON(http.StatusCreated, res)
}

// Update handler for updating blog
func Update(c echo.Context) error {
	var blog models.Blog
	if err := c.Bind(&blog); err != nil {
		res := response.New(response.MsgError, "invalid body to parser json", http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	// TODO validate zero values

	if blog.ID == 0 {
		res := response.New(response.MsgError, handlers.ErrIdRequired.Error(), http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	if err := psqlblog.Update(blog); err != nil {
		if errors.Is(err, storage.ErrNotFoundUpdate) {
			res := response.New(response.MsgError, err.Error(), http.StatusNotFound, nil)
			return c.JSON(http.StatusNotFound, res)
		}
		res := response.New(response.MsgError, err.Error(), http.StatusInternalServerError, nil)
		return c.JSON(http.StatusInternalServerError, res)
	}

	res := response.New(response.MsgOK, "Blog updated", http.StatusOK, blog)
	return c.JSON(http.StatusOK, res)
}

// Delete handler for deleting blog
func Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := response.New(response.MsgError, handlers.ErrIdRequired.Error(), http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	if err := psqlblog.Delete(uint(id)); err != nil {
		if errors.Is(err, storage.ErrNotFoundDelete) {
			res := response.New(response.MsgError, err.Error(), http.StatusNotFound, nil)
			return c.JSON(http.StatusNotFound, res)
		}
		res := response.New(response.MsgError, err.Error(), http.StatusInternalServerError, nil)
		return c.JSON(http.StatusInternalServerError, res)
	}

	res := response.New(response.MsgOK, "Blog deleted", http.StatusOK, nil)
	return c.JSON(http.StatusOK, res)
}
