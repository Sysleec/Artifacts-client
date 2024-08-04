package main

import (
	"github.com/Sysleec/Artifacts-client/internal/models"
	"github.com/Sysleec/Artifacts-client/internal/repl"
)

func main() {
	cfg := models.Config{}
	repl.Run(&cfg)
}
