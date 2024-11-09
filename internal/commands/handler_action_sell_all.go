package commands

import (
	"github.com/Sysleec/Artifacts-client/internal/models"
	"github.com/Sysleec/Artifacts-client/internal/utils"
)

func commandSellAll(cfg *models.Config, _ ...string) error {
	character, err := utils.CheckCharacter(cfg, cfg.ApiClient.Character)
	if err != nil {
		return err
	}

	err = utils.SellAllItemsAndReturnToSpot(cfg, character)
	if err != nil {
		return err
	}

	return nil
}
