package repl

import (
	"bufio"
	"fmt"
	"os"
)

func Run() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Artifacts > ")
		scanner.Scan()
		words := scanner.Text()
		if len(words) == 0 {
			continue
		}
	}
}
