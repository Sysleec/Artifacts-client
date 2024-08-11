package bot

import (
	"encoding/json"
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"github.com/Sysleec/Artifacts-client/internal/utils"
	"time"
)

func (c *ClientWrapper) Fighting(monster, char string) error {
	c.Client.BotRunning[char] = true

	var coords models.MoveReq

	switch monster {
	case "chicken":
		coords = models.MoveReq{
			X: 0,
			Y: 1,
		}
	case "cow":
		coords = models.MoveReq{
			X: 0,
			Y: 2,
		}
	case "green_slime":
		coords = models.MoveReq{
			X: 3,
			Y: -2,
		}
	case "yellow_slime":
		coords = models.MoveReq{
			X: 4,
			Y: -1,
		}
	case "red_slime":
		coords = models.MoveReq{
			X: 1,
			Y: -1,
		}
	case "blue_slime":
		coords = models.MoveReq{
			X: 2,
			Y: -1,
		}
	}

	err := c.goToSpot(coords, char)
	if err != nil {
		return fmt.Errorf("failed to go to %s: %w", monster, err)
	}

	go c.fight(char)

	return nil
}

func (c *ClientWrapper) goToSpot(coords models.MoveReq, char string) error {
	req := models.MoveReq{
		X: coords.X,
		Y: coords.Y,
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

	fmt.Printf("Moving to spot...\n")

	time.Sleep(time.Duration(action.Data.Cooldown.TotalSeconds) * time.Second)

	return nil
}

func (c *ClientWrapper) fight(char string) {
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
	fmt.Printf("\rFighting bot for character %s stopped			\n", char)
}
