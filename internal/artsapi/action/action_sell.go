package action

import (
	"encoding/json"
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"time"
)

func (c *ClientWrapper) Sell(req models.SellReq) (models.Transaction, error) {
	var actionMove models.Action

	act, err := c.Move(models.MoveReq{
		X: 5,
		Y: 1,
	})
	if err != nil {
		return models.Transaction{}, err
	}
	actionMove = act

	fmt.Printf("Moved character to x = %d, y = %d\n", actionMove.Data.Destination.X, actionMove.Data.Destination.Y)

	secondsRemaining := actionMove.Data.Cooldown.TotalSeconds
	fmt.Printf("Waiting for %d seconds\n", secondsRemaining)

	for secondsRemaining > 0 {
		fmt.Printf("\rTime left: %d seconds", secondsRemaining)
		time.Sleep(1 * time.Second)
		secondsRemaining--
	}

	fmt.Print("\rCooldown complete!              \n")

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
