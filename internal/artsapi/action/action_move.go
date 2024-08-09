package action

import (
	"encoding/json"
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func (c *ClientWrapper) Move(req models.MoveReq) (models.Action, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return models.Action{}, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.Client.PostReq("/my/"+c.Client.Character+"/action/move", body)
	if err != nil {
		return models.Action{}, fmt.Errorf("failed to send request: %w", err)
	}

	var action models.Action
	err = json.Unmarshal(resp, &action)
	if err != nil {
		return models.Action{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return action, nil
}
