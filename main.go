package main

import (
	"flag"
	"git.agehadev.com/elliebelly/jamboy/internal"
	"git.agehadev.com/elliebelly/jamboy/internal/engine"
	"go.uber.org/zap"
	"os"
)

var EmulationFinished = false

func main() {
	cartPath := flag.String("cart", "", "Path to cartridge")

	flag.Parse()

	if cartPath == nil || *cartPath == "" {
		internal.Logger.Error("cart path not provided")
		os.Exit(1)
	}

	cart := &engine.Cart{
		Path: *cartPath,
	}

	err := cart.Load()

	if err != nil {
		internal.Logger.Error(
			"failed to open cartridge",
			zap.Error(err),
			zap.String("cartridge", cart.Path),
		)
		os.Exit(1)
	}

	jamboy := engine.NewJamboy()

	jamboy.InsertCartridge(cart)
	jamboy.PowerOn()

	for !EmulationFinished {
		err := jamboy.Tick()

		if err != nil {
			internal.Logger.Error("we're in a bit of jam", zap.Error(err))
			EmulationFinished = true
		}
	}
}
