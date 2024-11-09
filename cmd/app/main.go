package main

import (
	"github.com/Sysleec/Artifacts-client/internal/artsapi"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"github.com/Sysleec/Artifacts-client/internal/repl"
	"github.com/Sysleec/Artifacts-client/internal/utils"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

const (
	httpTimeout = 5 * time.Second
	dbPath      = "./artifacts.db"
)

func main() {
	var accounts []models.Account

	DB, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open database")
	}

	tokens, err := utils.LoadTokens()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load tokens")
	}

	err = DB.AutoMigrate(&models.Account{})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to migrate database")
	}

	for name, token := range tokens {
		DB.FirstOrCreate(&models.Account{}, models.Account{Name: name, Token: token})
	}

	DB.Find(&accounts)
	if len(accounts) == 0 {
		log.Fatal().Err(err).Msg("No accounts found. Please add an account to config.ini")
	}

	defaultAcc := models.Account{}

	for _, acc := range accounts {
		if acc.IsDefault {
			defaultAcc = acc
			break
		}
	}

	apiClient := artsapi.NewClient(httpTimeout, defaultAcc.Token)

	cfg := models.Config{
		ApiClient: &apiClient,
		DB:        DB,
	}
	repl.Run(&cfg)
}
