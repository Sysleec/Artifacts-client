package characters

import (
	"encoding/json"
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func (c *ClientWrapper) Delete(request models.CharacterDeleteRequest) (models.CharacterResponse, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return models.CharacterResponse{}, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.Client.PostReq("/characters/delete", body)
	if err != nil {
		return models.CharacterResponse{}, fmt.Errorf("failed to send request: %w", err)
	}

	var character models.CharacterResponse
	err = json.Unmarshal(resp, &character)
	if err != nil {
		return models.CharacterResponse{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return character, nil
}
