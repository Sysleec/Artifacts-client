package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func commandHelp(_ *models.Config, _ ...string) error {
	fmt.Println("Welcome to the Artifacts client help!")
	fmt.Println("Available List:")
	for _, comm := range List() {
		fmt.Println(comm.name + " - " + comm.description)
	}
	return nil
}
