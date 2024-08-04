package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"os"
)

func commandExit(_ *models.Config, _ ...string) error {
	fmt.Println("closing the Artifacts client...")
	os.Exit(0)
	return nil
}
