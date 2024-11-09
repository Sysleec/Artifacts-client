package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
	dbmodels "github.com/Sysleec/Artifacts-client/internal/models/DB"
)

func commandAccountSet(cfg *models.Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected exactly  argument, got %d", len(args))
	}

	accountName := args[0]

	var account dbmodels.Account

	cfg.DB.Model(&dbmodels.Account{}).Update("is_default", false)

	if err := cfg.DB.Where("name = ?", accountName).First(&account).Error; err != nil {
		return fmt.Errorf("account %s not found", accountName)
	}

	account.IsDefault = true

	if err := cfg.DB.Save(&account).Error; err != nil {
		return fmt.Errorf("failed to save account %s as default", accountName)
	}

	cfg.ApiClient.SetToken(account.Token)

	fmt.Printf("AccountDB %s set as active account.\n", accountName)

	return nil
}
