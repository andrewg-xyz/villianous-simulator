package main

import (
	"github.com/andrewg-xyz/villianous-simulator/src/cmd"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	cmd.Execute()
}
