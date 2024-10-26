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
)

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())
	w := a.NewWindow("Artifacts GUI client")
	w.Resize(fyne.NewSize(500, 500))
	w.CenterOnScreen()

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

	newTaskIcon, err := fyne.LoadResourceFromPath("internal/icons/add.png")

	accountsBar := container.NewHBox(
		canvas.NewText(" Accounts", color.Black),
		layout.NewSpacer(),
		widget.NewButtonWithIcon(" Add", newTaskIcon, func() {
			log.Info().Msg("Add account")
		}),
	)

	//container := accounts.CheckOrCreateAccounts(DB)
	//if err != nil {
	//	return
	//}

	_ = DB

	accountContent := container.NewVBox(
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
