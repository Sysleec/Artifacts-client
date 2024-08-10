package action

import (
	"encoding/json"
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func (c *ClientWrapper) Gathering() (models.Action, error) {
	resp, err := c.Client.PostReq("/my/"+c.Client.Character+"/action/gathering", []byte{})
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
