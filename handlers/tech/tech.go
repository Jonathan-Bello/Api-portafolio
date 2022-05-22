package tech

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Jonathan-Bello/Api-portafolio/handlers"
	"github.com/Jonathan-Bello/Api-portafolio/models"
	"github.com/Jonathan-Bello/Api-portafolio/models/response"
	"github.com/Jonathan-Bello/Api-portafolio/storage"
	psqltech "github.com/Jonathan-Bello/Api-portafolio/storage/psql_tech"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// GetAll handler for getting all techs
func GetAll(c echo.Context) error {
	techs, err := psqltech.GetAll()
	if err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			res := response.New(response.MsgError, err.Error(), http.StatusBadRequest, nil)
			return c.JSON(http.StatusBadRequest, res)
		}

		res := response.New(response.MsgError, err.Error(), http.StatusInternalServerError, nil)
		return c.JSON(http.StatusInternalServerError, res)
	}

	res := response.New(response.MsgOK, "OK", http.StatusOK, techs)
	return c.JSON(http.StatusOK, res)
}

// GetByID handler for getting tech by id
func GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := response.New(response.MsgError, handlers.ErrIdRequired.Error(), http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	tech, err := psqltech.GetByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res := response.New(response.MsgError, err.Error(), http.StatusNotFound, nil)
			return c.JSON(http.StatusNotFound, res)
		}

		res := response.New(response.MsgError, err.Error(), http.StatusInternalServerError, nil)
		return c.JSON(http.StatusInternalServerError, res)
	}

	res := response.New(response.MsgOK, "OK", http.StatusOK, tech)
	return c.JSON(http.StatusOK, res)
}

// Create handler for creating tech
func Create(c echo.Context) error {
	var tech models.Tech
	if err := c.Bind(&tech); err != nil {
		res := response.New(response.MsgError, handlers.ErrInvalidBody.Error(), http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	//TODO: Validate tech data with zero value

	if err := psqltech.Create(tech); err != nil {
		res := response.New(response.MsgError, err.Error(), http.StatusInternalServerError, nil)
		return c.JSON(http.StatusInternalServerError, res)
	}

	res := response.New(response.MsgOK, "Tech created", http.StatusCreated, tech)
	return c.JSON(http.StatusCreated, res)
}

// Update handler for updating tech
func Update(c echo.Context) error {
	var tech models.Tech
	if err := c.Bind(&tech); err != nil {
		res := response.New(response.MsgError, handlers.ErrInvalidBody.Error(), http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	if tech.ID == 0 {
		res := response.New(response.MsgError, handlers.ErrIdRequired.Error(), http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	if err := psqltech.Update(tech); err != nil {
		if errors.Is(err, storage.ErrNotFoundUpdate) {
			res := response.New(response.MsgError, err.Error(), http.StatusNotFound, nil)
			return c.JSON(http.StatusNotFound, res)
		}
		res := response.New(response.MsgError, err.Error(), http.StatusInternalServerError, nil)
		return c.JSON(http.StatusInternalServerError, res)
	}

	res := response.New(response.MsgOK, "Tech updated", http.StatusOK, nil)
	return c.JSON(http.StatusOK, res)
}

// Delete handler for deleting tech
func Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := response.New(response.MsgError, handlers.ErrIdRequired.Error(), http.StatusBadRequest, nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	if err := psqltech.Delete(uint(id)); err != nil {
		if errors.Is(err, storage.ErrNotFoundDelete) {
			res := response.New(response.MsgError, err.Error(), http.StatusNotFound, nil)
			return c.JSON(http.StatusNotFound, res)
		}
		res := response.New(response.MsgError, err.Error(), http.StatusInternalServerError, nil)
		return c.JSON(http.StatusInternalServerError, res)
	}

	res := response.New(response.MsgOK, "Tech deleted", http.StatusOK, nil)
	return c.JSON(http.StatusOK, res)
}
