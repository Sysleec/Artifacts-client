package account

import (
	"encoding/json"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func (c *ClientWrapper) Get() (models.Account, error) {
	resp, err := c.Client.GetReq("/my/details")
	if err != nil {
		return models.Account{}, err
	}

	var account models.Account

	err = json.Unmarshal(resp, &account)
	if err != nil {
		return models.Account{}, err
	}

	return account, nil
}
