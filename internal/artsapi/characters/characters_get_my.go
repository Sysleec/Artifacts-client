package characters

import (
	"encoding/json"
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func (c *ClientWrapper) GetMyCharacters() (models.MyCharactersResponse, error) {
	resp, err := c.Client.GetReq("/my/characters")
	if err != nil {
		return models.MyCharactersResponse{}, fmt.Errorf("failed to send request: %w", err)
	}

	var characters models.MyCharactersResponse
	err = json.Unmarshal(resp, &characters)
	if err != nil {
		return models.MyCharactersResponse{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return characters, nil
}
