package utils

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func CheckMaxItems(character models.Character) bool {
	maxItemsQty := character.InventoryMaxItems
	itemsQty := 0

	for _, item := range character.Inventory {
		itemsQty += item.Quantity
	}

	if itemsQty+3 >= maxItemsQty {
		fmt.Printf("Character %s has reached the maximum number of items\n", character.Name)
		return true
	}

	return false
}
