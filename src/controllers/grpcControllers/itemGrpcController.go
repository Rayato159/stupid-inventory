package grpcControllers

import (
	"fmt"
	"io"
	"time"

	pbItem "github.com/Rayato159/stupid-inventory/src/proto/item"
	"github.com/Rayato159/stupid-inventory/src/repositories"
	"golang.org/x/net/context"
)

type ItemGrpcController struct {
	ItemRepository *repositories.ItemRepository
	pbItem.UnimplementedItemServiceServer
}

func (h *ItemGrpcController) FindItems(stream pbItem.ItemService_FindItemsServer) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	items := &pbItem.ItemArr{
		Data: make([]*pbItem.Item, 0),
	}

	for {
		itemId, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("item_id out of range")
			break
		}
		if err != nil {
			return err
		}

		item, err := h.ItemRepository.FindOneItem(ctx, itemId.Id)
		if err != nil {
			return err
		}
		items.Data = append(items.Data, &pbItem.Item{
			Id:          item.ObjectId.Hex(),
			Title:       item.Title,
			Description: item.Description,
			Damage:      item.Damage,
		})
	}

	stream.SendAndClose(items)
	return nil
}
