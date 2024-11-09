package commands

import (
	"encoding/json"
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
	dbmodels "github.com/Sysleec/Artifacts-client/internal/models/DB"
	"time"
)

func commandMapUpload(cfg *models.Config, _ ...string) error {
	fmt.Println("Uploading maps...")

	maxPages := 1

	for page := 1; page <= maxPages; page++ {
		time.Sleep(500 * time.Millisecond)
		maps := models.Maps{}

		req, err := cfg.ApiClient.GetReq(fmt.Sprintf("/maps?page=%d", page))
		if err != nil {
			return err
		}

		err = json.Unmarshal(req, &maps)
		if err != nil {
			return err
		}

		maxPages = maps.Pages

		for _, mapsData := range maps.Data {
			maps := dbmodels.Maps{
				Name: mapsData.Name,
				Skin: mapsData.Skin,
				X:    mapsData.X,
				Y:    mapsData.Y,
				Type: mapsData.Content.Type,
				Code: mapsData.Content.Code,
			}

			err = cfg.DB.FirstOrCreate(&maps, dbmodels.Maps{X: maps.X, Y: maps.Y}).Error
		}
	}

	fmt.Println("Maps uploaded")

	return nil
}
