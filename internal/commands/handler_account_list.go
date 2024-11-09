package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func commandAccountList(cfg *models.Config, args ...string) error {
	var accounts []models.Account

	cfg.DB.Find(&accounts)

	if len(accounts) == 0 {
		return fmt.Errorf("no accounts found. Please add an account to config.ini")
	}

	fmt.Println("Accounts:")

	for _, acc := range accounts {
		fmt.Printf("Account: %s (default: %t)\n", acc.Name, acc.IsDefault)
	}

	return nil
}
