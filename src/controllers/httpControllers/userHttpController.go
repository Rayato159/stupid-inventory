package httpControllers

import (
	"net/http"
	"strings"

	"github.com/Rayato159/stupid-inventory/src/models"
	"github.com/Rayato159/stupid-inventory/src/repositories"
	"github.com/labstack/echo/v4"
)

type UserHttpController struct {
	UserRepository *repositories.UserRepository
}

func (h *UserHttpController) FindOneUser(c echo.Context) error {
	ctx := c.Request().Context()

	userId := strings.Trim(c.Param("user_id"), " ")

	user, err := h.UserRepository.FindOneUser(ctx, userId)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			models.Error{
				Message: err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, user)
}
