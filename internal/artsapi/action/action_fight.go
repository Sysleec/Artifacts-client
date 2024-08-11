package action

import (
	"encoding/json"
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func (c *ClientWrapper) Fight() (models.Fight, error) {
	resp, err := c.Client.PostReq("/my/"+c.Client.Character+"/action/fight", []byte{})
	if err != nil {
		return models.Fight{}, fmt.Errorf("failed to send request: %w", err)
	}

	var fight models.Fight
	err = json.Unmarshal(resp, &fight)
	if err != nil {
		return models.Fight{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return fight, nil
}
