package utils

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
	dbmodels "github.com/Sysleec/Artifacts-client/internal/models/DB"
)

func GetMapContent(cfgApiClient *models.Config, code string) ([]models.ContentWithCoords, error) {
	var maps []dbmodels.Maps
	err := cfgApiClient.DB.Where("type IS NOT NULL AND code IS NOT NULL AND code = ?", code).Find(&maps).Error
	if err != nil {
		return nil, fmt.Errorf("error getting maps: %w", err)
	}

	mapsContentCoords := make([]models.ContentWithCoords, 0)

	for _, mapData := range maps {
		mapsContentCoords = append(mapsContentCoords, models.ContentWithCoords{
			Code: mapData.Code,
			X:    mapData.X,
			Y:    mapData.Y,
		})
	}

	return mapsContentCoords, nil
}
