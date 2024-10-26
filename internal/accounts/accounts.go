package accounts

import (
	"encoding/base64"
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name  string
	Token string `gorm:"unique"`
}

//func CheckOrCreateAccounts(DB *gorm.DB) *fyne.Container {
//	err := DB.AutoMigrate(&Account{})
//	if err != nil {
//		log.Error().Err(err).Msg("failed to migrate accounts")
//		return nil
//	}
//
//	return container.NewVBox(
//		welcomeMsg2,
//		welcomeMsg3,
//		nameText,
//		nameEntry,
//		tokenText,
//		tokenEntry,
//		btnLogin,
//	)
//}

func AddAccountButton(DB *gorm.DB, namePtr, tokenPtr *string) func() {
	return func() {
		name, token := *namePtr, *tokenPtr

		if len(token) == 0 {
			log.Info().Msg("Token is empty")
			return
		}

		tokenBase64 := base64.StdEncoding.EncodeToString([]byte(token))

		var account Account
		DB.Where("token = ?", tokenBase64).First(&account)
		if account.Token == tokenBase64 {
			log.Info().Msg("Account already exists")
			return
		}

		accName := name
		if len(accName) == 0 {
			DB.Last(&account)
			accName = fmt.Sprintf("Account %d", account.ID+1)
		}

		account = Account{
			Name:  accName,
			Token: tokenBase64,
		}
		DB.Create(&account)
		log.Info().Msg("Account added")
	}
}
