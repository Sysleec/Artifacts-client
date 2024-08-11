package action

import (
	"encoding/json"
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func (c *ClientWrapper) Equip(req models.EquipReq) (models.Equip, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return models.Equip{}, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.Client.PostReq("/my/"+c.Client.Character+"/action/equip", body)
	if err != nil {
		return models.Equip{}, fmt.Errorf("failed to send request: %w", err)
	}

	var equip models.Equip
	err = json.Unmarshal(resp, &equip)
	if err != nil {
		return models.Equip{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return equip, nil
}
