package engine

import (
	"fmt"
	"git.agehadev.com/elliebelly/jamboy/internal"
	"go.uber.org/zap"
	"reflect"
	"runtime"
	"strings"
)

type Jamboy struct {
	CPU       *CPU
	Cartridge *Cart
	Memory    *Memory
}

func (j *Jamboy) InsertCartridge(cart *Cart) {
	j.Cartridge = cart
}

func (j *Jamboy) PowerOn() {
	j.CPU.Reset()
	j.Memory.Reset()

	j.CPU.CurrentOP = j.NextOp()
}

func (j *Jamboy) Tick() error {
	op := j.CPU.CurrentOP

	internal.Logger.Info("exec op",
		zap.String("func", GetFunctionName((*j.CPU.CurrentJumpTable)[op])),
		zap.String("code", fmt.Sprintf("0x%x", uint8(op))),
		zap.String("PC", fmt.Sprintf("0x%x", j.CPU.PC - 1)),
	)

	_, err := (*j.CPU.CurrentJumpTable)[op](j, op)

	if err != nil {
		return err
	}

	j.CPU.CurrentOP = j.NextOp()

	return nil
}

func (j *Jamboy) NextOp() (op OpCode) {
	j.CPU.Wait(1)

	op = OpCode(j.Cartridge.Data[j.CPU.PC])

	j.CPU.PC += 1

	return
}

func (j *Jamboy) Read16Bit() uint16 {
	pcl := j.NextOp()
	pch := j.NextOp()

	internal.Logger.Info("read 16bit", zap.String("value", fmt.Sprintf("0x%x", (uint16(pch)<<8)|uint16(pcl))))

	return (uint16(pch) << 8) | uint16(pcl)
}

func (j *Jamboy) Read8Bit() uint8 {
	num := uint8(j.NextOp())

	internal.Logger.Info("read 8bit", zap.Uint8("value", num))

	return num
}

func NewJamboy() *Jamboy {
	memory := &Memory{}

	return &Jamboy{
		Memory: memory,
		CPU:    NewCPU(memory),
	}
}

func GetFunctionName(i interface{}) string {
	return strings.Split(runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name(), "/engine.")[1]
}
