package bot

import (
	"encoding/json"
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"time"
)

func (c *ClientWrapper) MiningCopper(char string) error {
	c.Client.BotRunning[char] = true

	err := c.goToCopper(char)
	if err != nil {
		return fmt.Errorf("failed to go to copper: %w", err)
	}

	go c.gatherCopper(char)

	return nil
}

func (c *ClientWrapper) goToCopper(char string) error {
	req := models.MoveReq{
		X: 2,
		Y: 0,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.Client.PostReq("/my/"+char+"/action/move", body)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	var action models.Action
	err = json.Unmarshal(resp, &action)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	fmt.Println("Moving to copper...")

	time.Sleep(time.Duration(action.Data.Cooldown.TotalSeconds) * time.Second)

	return nil
}

func (c *ClientWrapper) gatherCopper(char string) {
	for c.Client.BotRunning[char] {
		resp, err := c.Client.PostReq("/my/"+char+"/action/gathering", []byte{})
		if err != nil {
			fmt.Errorf("failed to send request: %w", err)
			return
		}

		var action models.Action
		err = json.Unmarshal(resp, &action)
		if err != nil {
			fmt.Errorf("failed to unmarshal response: %w", err)
			return
		}

		time.Sleep(time.Duration(action.Data.Cooldown.TotalSeconds) * time.Second)
	}
	fmt.Printf("\rMining copper for %s stopped              \n", char)
}
