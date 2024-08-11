package utils

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func CheckMaxItems(character models.Character) error {
	maxItemsQty := character.InventoryMaxItems
	itemsQty := 0

	for _, item := range character.Inventory {
		itemsQty += item.Quantity
	}

	if itemsQty >= maxItemsQty {
		return fmt.Errorf("max items reached: %d/%d \n", itemsQty, maxItemsQty)
	}

	return nil
}
