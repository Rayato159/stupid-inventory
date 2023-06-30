package httpControllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Rayato159/stupid-inventory/pkg/utils"
	"github.com/Rayato159/stupid-inventory/src/config"
	"github.com/Rayato159/stupid-inventory/src/models"
	pbItem "github.com/Rayato159/stupid-inventory/src/proto/item"
	"github.com/Rayato159/stupid-inventory/src/repositories"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserHttpController struct {
	Cfg            *config.Config
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

	// Pulling items on item app using gRPC
	connItem, err := grpc.Dial(h.Cfg.Grpc.ItemAppUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.Error{
			Message: err.Error(),
		})
	}
	defer connItem.Close()

	clientItem := pbItem.NewItemServiceClient(connItem)

	streamItems, err := clientItem.FindItems(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.Error{
			Message: err.Error(),
		})
	}
	for _, item := range user.Items {
		if err := streamItems.Send(
			&pbItem.ItemReq{
				Id: item.ObjectId.Hex(),
			},
		); err != nil {
			log.Printf("%v.Send(%v) = %v", streamItems, &pbItem.ItemReq{Id: item.ObjectId.Hex()}, err)
			return c.JSON(http.StatusInternalServerError, &models.Error{
				Message: err.Error(),
			})
		}
		fmt.Printf("order: %v has been streamed\n", &pbItem.ItemReq{Id: item.ObjectId.Hex()})
	}
	results, err := streamItems.CloseAndRecv()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.Error{
			Message: err.Error(),
		})
	}
	if len(results.Data) == 0 {
		return c.JSON(http.StatusInternalServerError, &models.Error{
			Message: err.Error(),
		})
	}

	user.Items = make([]models.Item, 0)
	for _, item := range results.Data {
		user.Items = append(user.Items, models.Item{
			ObjectId:    utils.BsonObjectID(item.Id),
			Title:       item.Title,
			Description: item.Description,
			Damage:      item.Damage,
		})
	}

	return c.JSON(http.StatusOK, user)
}
