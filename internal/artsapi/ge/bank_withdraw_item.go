package ge

import (
	"encoding/json"
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func (c ClientWrapper) WithdrawItem(req models.BankReq) (models.BankTransaction, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return models.BankTransaction{}, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.Client.PostReq("/my/"+c.Client.Character+"/action/bank/withdraw", body)
	if err != nil {
		return models.BankTransaction{}, fmt.Errorf("failed to send request: %w", err)
	}

	var bankTransaction models.BankTransaction
	err = json.Unmarshal(resp, &bankTransaction)
	if err != nil {
		return models.BankTransaction{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return bankTransaction, nil
}
