package ge

import (
	"encoding/json"
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func (c *ClientWrapper) GetItem(item string) (models.GeItem, error) {
	resp, err := c.Client.GetReq(fmt.Sprintf("/grandexchange/orders?code=%s", item))
	if err != nil {
		return models.GeItem{}, fmt.Errorf("failed to send request: %w", err)
	}

	var geItem models.GeItem
	err = json.Unmarshal(resp, &geItem)
	if err != nil {
		return models.GeItem{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return geItem, nil
}
