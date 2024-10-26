package models

import (
	"github.com/Sysleec/Artifacts-client/internal/artsapi"
	"gorm.io/gorm"
)

type Config struct {
	ApiClient *artsapi.Client
	DB        *gorm.DB
}
