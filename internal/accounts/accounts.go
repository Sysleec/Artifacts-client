package accounts

import (
	"encoding/base64"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"net/url"
)

type Account struct {
	gorm.Model
	Name  string
	Token string `gorm:"unique"`
}

func CheckOrCreateAccounts(DB *gorm.DB) *fyne.Container {
	err := DB.AutoMigrate(&Account{})
	if err != nil {
		log.Error().Err(err).Msg("failed to migrate accounts")
		return nil
	}

	welcomeMsg2 := widget.NewLabel("You need to enter your token to start using the client.")
	welcomeMsg3 := widget.NewHyperlink("You can get it here", &url.URL{Scheme: "https", Host: "artifactsmmo.com/account"})
	nameText := widget.NewLabel("Account name (optional)")
	nameEntry := widget.NewEntry()
	tokenText := widget.NewLabel("Token")
	tokenEntry := widget.NewPasswordEntry()
	btnLogin := widget.NewButton("Add account", func() {
		token := tokenEntry.Text
		if len(token) == 0 {
			return
		}

		tokenBase64 := base64.StdEncoding.EncodeToString([]byte(token))

		var account Account
		DB.Where("token = ?", tokenBase64).First(&account)
		if account.Token == tokenBase64 {
			log.Info().Msg("logged in")
			return
		}

		accName := nameEntry.Text
		if len(accName) == 0 {
			DB.Last(&account)
			accName = fmt.Sprintf("Account %d", account.ID+1)
		}

		account = Account{
			Name:  accName,
			Token: tokenBase64,
		}
		DB.Create(&account)
	})

	return container.NewVBox(
		welcomeMsg2,
		welcomeMsg3,
		nameText,
		nameEntry,
		tokenText,
		tokenEntry,
		btnLogin,
	)
}
