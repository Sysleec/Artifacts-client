package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"github.com/Sysleec/Artifacts-client/internal/utils"
)

func commandMapGet(cfg *models.Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected exactly 1 argument, got %d", len(args))
	}

	fmt.Println("Getting map content...")

	content, err := utils.GetMapContent(cfg, args[0])
	if err != nil {
		return err
	}

	for _, c := range content {
		fmt.Printf("X: %d, Y: %d, Code: %s\n", c.X, c.Y, c.Code)
	}

	return nil
}
