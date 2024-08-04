package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"os"
)

func commandExit(cfg *models.Config, args ...string) error {
	fmt.Println("closing the Artifacts client...")
	os.Exit(0)
	return nil
}
