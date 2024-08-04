package main

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"github.com/Sysleec/Artifacts-client/internal/repl"
	"github.com/Sysleec/Artifacts-client/internal/utils"
	"time"
)

func main() {
	fmt.Println("Welcome to the Artifacts client!")
	fmt.Println("Type 'help' for a list of available commands")

	tok, err := utils.LoadToken()
	if err != nil {
		fmt.Println(err)
		return
	}

	apiClient := artsapi.NewClient(60*time.Second, tok)

	cfg := models.Config{
		ApiClient: &apiClient,
	}
	repl.Run(&cfg)
}
