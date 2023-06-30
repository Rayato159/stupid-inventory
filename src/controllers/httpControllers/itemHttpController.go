package httpControllers

import (
	"net/http"
	"strings"

	"github.com/Rayato159/stupid-inventory/src/models"
	"github.com/Rayato159/stupid-inventory/src/repositories"
	"github.com/labstack/echo/v4"
)

type ItemHttpController struct {
	ItemRepository *repositories.ItemRepository
}

func (h *ItemHttpController) FindItems(c echo.Context) error {
	ctx := c.Request().Context()

	items := make([]models.Item, 0)
	if err := h.ItemRepository.FindItems(ctx, &items); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			models.Error{
				Message: err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, items)
}

func (h *ItemHttpController) FindOneItem(c echo.Context) error {
	ctx := c.Request().Context()

	itemId := strings.Trim(c.Param("item_id"), " ")

	item, err := h.ItemRepository.FindOneItem(ctx, itemId)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			models.Error{
				Message: err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, item)
}
