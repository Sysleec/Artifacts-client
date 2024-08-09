package repl

import (
	"bufio"
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/commands"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"os"
	"strings"
)

func Run(cfg *models.Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Artifacts > ")
		scanner.Scan()
		words := scanner.Text()
		if len(words) == 0 {
			continue
		}

		wordsSl := commFormatter(words)

		if len(wordsSl) > 4 {
			fmt.Println("Too many args")
			continue
		}

		var commandName string
		var args []string

		// Find the command in the command list
		for i := len(wordsSl); i > 0; i-- {
			commandName = strings.Join(wordsSl[:i], " ")
			if comm, ok := commands.List()[commandName]; ok {
				args = wordsSl[i:]
				err := comm.Callback(cfg, args...)
				if err != nil {
					fmt.Println(err)
				}
				break
			}
			// If command not found
			commandName = ""
		}

		if commandName == "" {
			fmt.Println("Command not found")
		}
	}
}

func commFormatter(words string) []string {
	if !strings.Contains(words, "get") {
		words = strings.ToLower(words)
	}
	wordsSlice := strings.Fields(words)

	return wordsSlice
}
