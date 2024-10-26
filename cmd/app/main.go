package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Sysleec/Artifacts-client/internal/accounts"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"image/color"
	"net/url"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())
	w := a.NewWindow("Artifacts GUI client")
	w.Resize(fyne.NewSize(500, 500))
	w.CenterOnScreen()

	var addAccountContent *fyne.Container
	var accountContent *fyne.Container
	var accs []accounts.Account

	DB, err := gorm.Open(sqlite.Open("artifacts.db"), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect database")
	}
	DB.Find(&accs)

	noAccountsLabel := canvas.NewText("No accounts found", color.Black)
	if len(accs) == 0 {
		noAccountsLabel.Hide()
	}

	accCreateMsg1 := widget.NewLabel("You need to enter your token to start using the client.")
	accCreateMsg2 := widget.NewHyperlink("You can get it here", &url.URL{Scheme: "https", Host: "artifactsmmo.com/account"})
	nameText := widget.NewLabel("Account name (optional)")
	nameEntry := widget.NewEntry()
	tokenText := widget.NewLabel("Token")
	tokenEntry := widget.NewPasswordEntry()
	btnLogin := widget.NewButton("Add account", accounts.AddAccountButton(DB, &nameEntry.Text, &tokenEntry.Text))

	newAccIcon, err := fyne.LoadResourceFromPath("internal/icons/add.png")
	backIcon, err := fyne.LoadResourceFromPath("internal/icons/back.png")

	accountsBar := container.NewHBox(
		canvas.NewText(" Accounts", color.Black),
		layout.NewSpacer(),
		widget.NewButtonWithIcon(" Add", newAccIcon, func() {
			w.SetContent(addAccountContent)
		}),
	)

	addAccountBar := container.NewHBox(
		canvas.NewText(" Add account", color.Black),
		layout.NewSpacer(),
		widget.NewButtonWithIcon(" Back", backIcon, func() {
			//nameEntry.SetText("")
			//tokenEntry.SetText("")
			//nameEntry.Refresh()
			//tokenEntry.Refresh()

			w.SetContent(accountContent)
		}),
	)

	addAccountContent = container.NewVBox(
		addAccountBar,
		canvas.NewLine(color.Black),
		accCreateMsg1,
		accCreateMsg2,
		nameText,
		nameEntry,
		tokenText,
		tokenEntry,
		btnLogin,
	)

	//_ = accounts.CheckOrCreateAccounts(DB)
	//if err != nil {
	//	return
	//}

	_ = DB

	accountContent = container.NewVBox(
		accountsBar,
		canvas.NewLine(color.Black),
		noAccountsLabel,
	)

	w.SetContent(accountContent)
	w.Show()
	a.Run()

	//apiClient := artsapi.NewClient(60*time.Second, createAccounts[0].Token)
	//
	//cfg := models.Config{
	//	ApiClient: &apiClient,
	//	DB:        DB,
	//}
	//repl.Run(&cfg)
}

//func main() {
//	fmt.Println("Welcome to the Artifacts client!")
//	fmt.Println("Type 'help' for a list of available commands")
//
//	tok, err := utils.LoadToken()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//}
