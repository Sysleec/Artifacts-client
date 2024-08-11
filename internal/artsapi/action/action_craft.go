package action

import (
	"encoding/json"
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"time"
)

func (c *ClientWrapper) Craft(req models.CraftReq) (models.Action, error) {
	var actionMove models.Action

	switch req.Code {
	case "copper":
		act, err := c.Move(models.MoveReq{
			X: 1,
			Y: 5,
		})
		if err != nil {
			return models.Action{}, err
		}

		actionMove = act
	default:
		return models.Action{}, fmt.Errorf("unknown target: %s", req.Code)
	}

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
		return models.Action{}, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.Client.PostReq("/my/"+c.Client.Character+"/action/crafting", body)
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
