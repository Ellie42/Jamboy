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
	"runtime"
	"strings"
	"unsafe"
)

var EmulationFinished = false

func main() {
	cartPath := flag.String("cart", "", "Path to cartridge")
	dump := flag.String("dump-at", "0x0000", "Dump all data at PC Xh")
	outputDebug := flag.Bool("output-debug", false, "Output debug text")
	bootROMPath := flag.String("boot-rom-path", "roms/dmg_boot.bin", "Path to DMG boot ROM")
	loopBoot := flag.Bool("loop-boot", false, "Loop loading the boot ROM")

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

	runtime.LockOSThread()

	testPixels := make([]uint8, engine.ResolutionX*engine.ResolutionY*4)

	window := renderer.NewWindow()
	jamboy := engine.NewJamboy()

	jamboy.GPU.PixelBuffer = &testPixels

	var done chan bool

	go func() {
		<-window.Initialised

		go runJamboy(jamboy, outputDebug, bootROMPath, cart, loopBoot, dump, dumpLine, done)
	}()

	window.Open(engine.ResolutionX, engine.ResolutionY, 4, unsafe.Pointer(&testPixels[0]))

	jamboy.PowerOff()
	EmulationFinished = true
}

func runJamboy(jamboy *engine.Jamboy, outputDebug *bool, bootROMPath *string, cart *engine.Cart, loopBoot *bool, dump *string, dumpLine uint16, done chan bool) {
	if outputDebug != nil {
		jamboy.OutputDebug = *outputDebug
	}

	var bootROMdata []byte = nil

	if bootROMPath != nil && *bootROMPath != "" {
		var err error
		bootROMdata, err = ioutil.ReadFile(*bootROMPath)

		if err != nil {
			panic(err)
		}
	}

	jamboy.InsertCartridge(cart)
	jamboy.PowerOn(bootROMdata)

	for !EmulationFinished {
		if *loopBoot && jamboy.CPU.PC > 0x100 {
			jamboy.PowerOn(jamboy.CPU.BootROM)
		}

		err := jamboy.Tick()

		if err != nil {
			internal.Logger.Error("we're in a bit of jam", zap.Error(err))
			EmulationFinished = true
		}

		if dump != nil && dumpLine > 0 && jamboy.CurrentOPAddr == engine.Address(dumpLine) {
			err := ioutil.WriteFile(
				fmt.Sprintf("dumps/jamboy_ram_dump_%04x.bin", jamboy.CurrentOPAddr),
				jamboy.MMU.RAM[:], 0777,
			)

			if err != nil {
				panic(err)
			}

			err = ioutil.WriteFile(
				fmt.Sprintf("dumps/jamboy_register_dump_%04x.txt", jamboy.CurrentOPAddr),
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
	}

	done <- true
}
