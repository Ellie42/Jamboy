package engine

import (
	"fmt"
	"git.agehadev.com/elliebelly/jamboy/internal"
	"reflect"
	"runtime"
	"strings"
)

type Jamboy struct {
	CPU      *CPU
	GPU      *GPU
	Cart     *internal.Cart
	MMU      *MMU
	Code     *Code
	Debugger *Debugger

	OutputDebug   bool
	IsHalted      bool
	BootROMPath   string
	CurrentOPAddr Address
	LoopBoot      bool
}

func (j *Jamboy) InsertCartridge(cart *internal.Cart) {
	j.Cart = cart
}

func (j *Jamboy) PowerOn(bootROMdata []byte) {
	if bootROMdata != nil {
		j.CPU.HasBootROM = true
	}

	j.CPU.Reset()
	j.MMU.Reset()
	j.GPU.Reset()
	j.Code.Reset()
	j.Debugger.Reset()

	if bootROMdata != nil {
		j.CPU.LoadBootRom(bootROMdata)
	}

	j.CPU.CurrentOP = j.NextOp()
}

func (j *Jamboy) Tick() error {
	j.Debugger.PreTick()

	if j.IsHalted {
		return nil
	}

	//if j.CurrentOPAddr == 0x00A3 {
	//	fmt.Printf("%x\n", j.CurrentOPAddr)
	//}

	op := j.CPU.CurrentOP

	opFunc := (*j.CPU.CurrentJumpTable)[op]

	if j.OutputDebug {
		fmt.Printf("%s: %x", GetFunctionName(opFunc), op)
	}

	j.CPU.CurrentJumpTable = &BaseOpJumpTable

	err := opFunc(j, op)

	if j.OutputDebug {
		fmt.Printf(`
AF %02x%02x BC %02x%02x
DE %02x%02x HL %02x%02x
SP %04x PC %04x
-------------------
`,
			j.CPU.Registers[A], j.CPU.Registers[F],
			j.CPU.Registers[B], j.CPU.Registers[C],
			j.CPU.Registers[D], j.CPU.Registers[E],
			j.CPU.Registers[H], j.CPU.Registers[L],
			j.CPU.SP, j.CurrentOPAddr,
		)
	}

	if err != nil {
		return err
	}

	j.GPU.Tick()

	if j.CPU.IsBooting && j.MMU.Read(AddrBootROMDisable) > 0 {
		j.CPU.UnloadBootROM()
	}

	j.Debugger.PostTick()

	j.CPU.CurrentOP = j.NextOpInstant()

	return nil
}

func (j *Jamboy) NextOpInstant() (op OpCode) {
	op = OpCode(j.MMU.ReadInstant(Address(j.CPU.PC)))

	j.CurrentOPAddr = Address(j.CPU.PC)

	j.CPU.PC += 1

	return
}

func (j *Jamboy) NextOp() (op OpCode) {
	j.CPU.Wait(1)

	op = OpCode(j.MMU.Read(Address(j.CPU.PC)))

	j.CPU.PC += 1

	return
}

func (j *Jamboy) Read16Bit() uint16 {
	pcl := j.NextOp()
	pch := j.NextOp()

	val := (uint16(pch) << 8) | uint16(pcl)

	if j.OutputDebug {
		fmt.Printf(" %x(d16)", val)
	}

	return val
}

func (j *Jamboy) Read8Bit() uint8 {
	num := uint8(j.NextOp())

	if j.OutputDebug {
		fmt.Printf(" %x(d8)", num)
	}

	return num
}

func (j *Jamboy) Halt() {
	j.IsHalted = true
}

func (j *Jamboy) Resume() {
	j.IsHalted = false
}

func (j *Jamboy) PowerOff() {
	j.IsHalted = true
}

func NewJamboy() *Jamboy {
	jb := &Jamboy{}

	jb.MMU = newMMU(jb)
	jb.CPU = NewCPU(jb)
	jb.GPU = NewGPU(jb)
	jb.Code = NewCode(jb)
	jb.Debugger = NewDebugger(jb)

	return jb
}

func GetFunctionName(i interface{}) string {
	return strings.Split(runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name(), "/engine.")[1]
}
