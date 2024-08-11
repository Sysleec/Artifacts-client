package utils

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func CheckMaxItems(action models.Action) error {
	maxItemsQty := action.Data.Character.InventoryMaxItems
	itemsQty := 0

	for _, item := range action.Data.Character.Inventory {
		itemsQty += item.Quantity
	}

	if itemsQty >= maxItemsQty {
		return fmt.Errorf("max items reached: %d/%d", itemsQty, maxItemsQty)
	}

	return nil
}
