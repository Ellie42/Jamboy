package engine

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

	return (uint16(pch) << 8) | uint16(pcl)
}

func (j *Jamboy) Read8Bit() uint {
	return uint(j.NextOp())
}

func NewJamboy() *Jamboy {
	memory := &Memory{}

	return &Jamboy{
		Memory: memory,
		CPU:    NewCPU(memory),
	}
}
