package commands

import (
	"encoding/json"
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
	dbmodels "github.com/Sysleec/Artifacts-client/internal/models/DB"
)

func commandResourcesUpload(cfg *models.Config, _ ...string) error {
	resources := models.Resources{}

	req, err := cfg.ApiClient.GetReq("/resources")
	if err != nil {
		return err
	}

	err = json.Unmarshal(req, &resources)
	if err != nil {
		return err
	}

	fmt.Println("Uploading resources...")

	for _, resource := range resources.Data {
		dbResource := dbmodels.Resource{
			Name:  resource.Name,
			Code:  resource.Code,
			Skill: resource.Skill,
			Level: resource.Level,
		}

		err = cfg.DB.Where("code = ?", resource.Code).FirstOrCreate(&dbResource).Error
		if err != nil {
			return err
		}

		var drops []dbmodels.Drop
		for _, drop := range resource.Drops {
			drops = append(drops, dbmodels.Drop{
				ResourceID:  dbResource.ID,
				Code:        drop.Code,
				Rate:        drop.Rate,
				MinQuantity: drop.MinQuantity,
				MaxQuantity: drop.MaxQuantity,
			})
		}

		err = cfg.DB.Where("resource_id = ?", dbResource.ID).Delete(&dbmodels.Drop{}).Error
		if err != nil {
			return err
		}
		err = cfg.DB.Create(&drops).Error
		if err != nil {
			return err
		}
	}

	fmt.Println("Resources uploaded")

	return nil
}
