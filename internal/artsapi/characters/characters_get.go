package characters

import (
	"encoding/json"
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func (c *ClientWrapper) Get(name string) (models.CharacterResponse, error) {
	resp, err := c.Client.GetReq("/characters/" + name)
	if err != nil {
		return models.CharacterResponse{}, fmt.Errorf("failed to send request: %w", err)
	}

	var characters models.CharacterResponse
	err = json.Unmarshal(resp, &characters)
	if err != nil {
		return models.CharacterResponse{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return characters, nil
}
