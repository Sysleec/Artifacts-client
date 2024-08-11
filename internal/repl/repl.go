package repl

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/commands"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"github.com/c-bata/go-prompt"
	"regexp"
	"strings"
)

func Run(cfg *models.Config) {
	p := prompt.New(
		executor(cfg),
		emptyCompleter,
		prompt.OptionPrefix("Artifacts > "),
		prompt.OptionTitle("Artifacts CLI"),
	)
	p.Run()
}

func executor(cfg *models.Config) func(string) {
	return func(input string) {
		input = strings.TrimSpace(input)
		if len(input) == 0 {
			return
		}

		wordsSl := commFormatter(input)

		if len(wordsSl) > 5 {
			fmt.Println("Too many args")
			return
		}

		var commandName string
		var args []string

		// Find command
		for i := len(wordsSl); i > 0; i-- {
			commandName = strings.Join(wordsSl[:i], " ")
			if comm, ok := commands.List()[commandName]; ok {
				args = wordsSl[i:]
				err := comm.Callback(cfg, args...)
				if err != nil {
					fmt.Println("Error:", err)
				}
				return
			}
		}

		fmt.Println("Command not found")
	}
}

func emptyCompleter(_ prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

func commFormatter(words string) []string {
	if !strings.Contains(words, "get") {
		words = strings.ToLower(words)
	}
	return splitQuotedWords(words)
}

func splitQuotedWords(s string) []string {
	re := regexp.MustCompile(`\"[^\"]+\"|\S+`)
	matches := re.FindAllString(s, -1)

	for i, match := range matches {
		if len(match) > 0 && match[0] == '"' {
			match = match[1 : len(match)-1]
			matches[i] = match
		}
	}

	return matches
}
