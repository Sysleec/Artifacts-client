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

		if len(wordsSl) > 2 {
			fmt.Println("Too many args")
			continue
		}

		commandName := wordsSl[0]
		args := []string{}
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
	lower := strings.ToLower(words)
	wordsSlice := strings.Fields(lower)

	return wordsSlice
}
