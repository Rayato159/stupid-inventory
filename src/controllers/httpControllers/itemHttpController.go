package httpControllers

import (
	"github.com/Rayato159/stupid-inventory/src/repositories"
	"github.com/labstack/echo/v4"
)

type ItemHttpController struct {
	ItemRepository *repositories.ItemRepository
}

func (h *ItemHttpController) FindItems(c echo.Context) error {
	return nil
}

func (h *ItemHttpController) FindOneItem(c echo.Context) error {
	return nil
}
