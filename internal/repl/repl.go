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

		if len(wordsSl) > 3 {
			fmt.Println("Too many args")
			continue
		}

		commandName := wordsSl[0]

		var args []string

		if len(wordsSl) > 1 {
			args = wordsSl[1:]
		}

		comm, ok := commands.List()[commandName]
		if ok {
			err := comm.Callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
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
