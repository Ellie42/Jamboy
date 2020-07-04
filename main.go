package main

import (
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"git.agehadev.com/elliebelly/jamboy/internal"
	"git.agehadev.com/elliebelly/jamboy/internal/engine"
	"git.agehadev.com/elliebelly/jamboy/internal/renderer"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"strings"
)

var EmulationFinished = false

func main() {
	cartPath := flag.String("cart", "", "Path to cartridge")
	dump := flag.String("dump-at", "0x0000", "Dump all data at PC Xh")
	outputDebug := flag.Bool("output-debug", false, "Output debug text")
	bootROMPath := flag.String("boot-rom-path", "roms/dmg_boot.bin", "Path to DMG boot ROM")
	stopAfterBoot := flag.Bool("stop-after-boot", false, "Stop execution after boot ROM run")

	flag.Parse()

	dumpLine := uint16(0)

	if dump != nil {
		*dump = strings.Replace(*dump, "0x", "", 1)
		hexBytes, err := hex.DecodeString(*dump)
		dumpLine = binary.BigEndian.Uint16(hexBytes)

		if err != nil {
			panic(err)
		}
	}

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

	window := renderer.NewWindow()
	jamboy := engine.NewJamboy()

	var done chan bool

	go func() {
		<-window.Initialised

		go runJamboy(jamboy, outputDebug, bootROMPath, cart, stopAfterBoot, dump, dumpLine, done)
	}()

	window.Open(engine.ResolutionX, engine.ResolutionY)

	jamboy.PowerOff()
	EmulationFinished = true
}

func runJamboy(jamboy *engine.Jamboy, outputDebug *bool, bootROMPath *string, cart *engine.Cart, stopAfterBoot *bool, dump *string, dumpLine uint16, done chan bool) {
	if outputDebug != nil {
		jamboy.OutputDebug = *outputDebug
	}

	if bootROMPath != nil && *bootROMPath != "" {
		rom, err := ioutil.ReadFile(*bootROMPath)

		if err != nil {
			panic(err)
		}

		jamboy.CPU.LoadBootRom(rom)
	}

	jamboy.InsertCartridge(cart)
	jamboy.PowerOn()

	for !EmulationFinished {
		if *stopAfterBoot && jamboy.CPU.PC > 0x100 {
			jamboy.PowerOn()
		}

		if dump != nil && dumpLine > 0 && jamboy.CPU.PC == dumpLine {
			err := ioutil.WriteFile(
				fmt.Sprintf("dumps/jamboy_ram_dump_%04x.bin", jamboy.CPU.PC),
				jamboy.MMU.RAM[:], 0777,
			)

			if err != nil {
				panic(err)
			}

			err = ioutil.WriteFile(
				fmt.Sprintf("dumps/jamboy_register_dump_%04x.bin", jamboy.CPU.PC),
				[]byte(fmt.Sprintf(`AF %02x%02x BC %02x%02x
DE %02x%02x HL %02x%02x
SP %04x PC %04x
-------------------
`,
					jamboy.CPU.Registers[engine.A], jamboy.CPU.Registers[engine.F],
					jamboy.CPU.Registers[engine.B], jamboy.CPU.Registers[engine.C],
					jamboy.CPU.Registers[engine.D], jamboy.CPU.Registers[engine.E],
					jamboy.CPU.Registers[engine.H], jamboy.CPU.Registers[engine.L],
					jamboy.CPU.SP, jamboy.CPU.PC,
				)), 0777,
			)

			if err != nil {
				panic(err)
			}

			break
		}

		err := jamboy.Tick()

		if err != nil {
			internal.Logger.Error("we're in a bit of jam", zap.Error(err))
			EmulationFinished = true
		}
	}

	done <- true
}
