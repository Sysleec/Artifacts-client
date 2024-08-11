package action

import (
	"encoding/json"
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func (c *ClientWrapper) Sell(req models.SellReq) (models.Transaction, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return models.Transaction{}, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.Client.PostReq("/my/"+c.Client.Character+"/action/ge/sell", body)
	if err != nil {
		return models.Transaction{}, fmt.Errorf("failed to send request: %w", err)
	}

	var trans models.Transaction
	err = json.Unmarshal(resp, &trans)
	if err != nil {
		return models.Transaction{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return trans, nil
}
