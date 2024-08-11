package bot

import (
	"encoding/json"
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"github.com/Sysleec/Artifacts-client/internal/utils"
	"time"
)

func (c *ClientWrapper) FightingChicken(char string) error {
	c.Client.BotRunning[char] = true

	err := c.goToChicken(char)
	if err != nil {
		return fmt.Errorf("failed to go to chicken: %w", err)
	}

	go c.fightChicken(char)

	return nil
}

func (c *ClientWrapper) goToChicken(char string) error {
	req := models.MoveReq{
		X: 0,
		Y: 1,
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

	fmt.Println("Moving to chickens...")

	time.Sleep(time.Duration(action.Data.Cooldown.TotalSeconds) * time.Second)

	return nil
}

func (c *ClientWrapper) fightChicken(char string) {
	for c.Client.BotRunning[char] {
		resp, err := c.Client.PostReq("/my/"+char+"/action/fight", []byte{})
		if err != nil {
			fmt.Printf("failed to send request: %s", err.Error())
			return
		}

		var action models.Action
		err = json.Unmarshal(resp, &action)
		if err != nil {
			fmt.Printf("failed to unmarshal response: %s", err.Error())
			return
		}

		err = utils.CheckMaxItems(models.ConvertToModelCharacter(action))
		if err != nil {
			fmt.Printf("\r%s\n", err.Error())
			return
		}

		time.Sleep(time.Duration(action.Data.Cooldown.TotalSeconds) * time.Second)
	}
	fmt.Printf("\rFighting chicken bot for character %s stopped			\n", char)
}