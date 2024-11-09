package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/account"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func commandAccountDetail(cfg *models.Config, _ ...string) error {
	client, err := account.NewClientWrapper(cfg.ApiClient)
	if err != nil {
		return fmt.Errorf("failed to action : %w", err)
	}

	acc, err := client.Get()
	if err != nil {
		return fmt.Errorf("failed to get account: %w", err)
	}

	fmt.Printf("Account Details: %+v\n", acc.Data)

	return nil
}
