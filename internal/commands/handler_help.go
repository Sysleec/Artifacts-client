package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func commandHelp(cfg *models.Config, args ...string) error {
	fmt.Println("Welcome to the Artifacts client help!")
	fmt.Println("Available List:")
	for _, comm := range List() {
		fmt.Println(comm.name + " - " + comm.description)
	}
	return nil
}