package main

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"github.com/Sysleec/Artifacts-client/internal/repl"
	"github.com/Sysleec/Artifacts-client/internal/utils"
)

func main() {
	fmt.Println("Welcome to the Artifacts client!")
	fmt.Println("Type 'help' for a list of available commands")

	tok, err := utils.LoadToken()
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg := models.Config{
		Token: tok,
	}
	repl.Run(&cfg)
}
