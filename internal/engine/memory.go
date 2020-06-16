package engine

type MemoryRange struct {
	From Address
	To   Address
}

var (
	MemoryROMLow = MemoryRange{
		From: Address(0x0000),
		To:   Address(0x3FFF),
	}
	MemoryROMHigh = MemoryRange{
		From: Address(0x3FFF),
		To:   Address(0x7FFF),
	}
)

type Address uint16

func (a Address) InRange(r MemoryRange) bool {
	if a >= r.From && a < r.To {
		return true
	}

	return false
}

//  0000-3FFF   16KB ROM Bank 00     (in cartridge, fixed at bank 00)
//  4000-7FFF   16KB ROM Bank 01..NN (in cartridge, switchable bank number)
//  8000-9FFF   8KB Video RAM (VRAM) (switchable bank 0-1 in CGB Mode)
//  A000-BFFF   8KB External RAM     (in cartridge, switchable bank, if any)
//  C000-CFFF   4KB Work RAM Bank 0 (WRAM)
//  D000-DFFF   4KB Work RAM Bank 1 (WRAM)  (switchable bank 1-7 in CGB Mode)
//  E000-FDFF   Same as C000-DDFF (ECHO)    (typically not used)
//  FE00-FE9F   Sprite Attribute Table (OAM)
//  FEA0-FEFF   Not Usable
//  FF00-FF7F   I/O Ports
//  FF80-FFFE   High RAM (HRAM)
//  FFFF        Interrupt Enable Register
type Memory struct {
	RAM [0x10000]byte
}

func (m *Memory) Reset() {
	m.RAM[0xFF05] = 0x00 //TIMA
	m.RAM[0xFF06] = 0x00 //TMA
	m.RAM[0xFF07] = 0x00 //TAC
	m.RAM[0xFF10] = 0x80 //NR10
	m.RAM[0xFF11] = 0xBF //NR11
	m.RAM[0xFF12] = 0xF3 //NR12
	m.RAM[0xFF14] = 0xBF //NR14
	m.RAM[0xFF16] = 0x3F //NR21
	m.RAM[0xFF17] = 0x00 //NR22
	m.RAM[0xFF19] = 0xBF //NR24
	m.RAM[0xFF1A] = 0x7F //NR30
	m.RAM[0xFF1B] = 0xFF //NR31
	m.RAM[0xFF1C] = 0x9F //NR32
	m.RAM[0xFF1E] = 0xBF //NR33
	m.RAM[0xFF20] = 0xFF //NR41
	m.RAM[0xFF21] = 0x00 //NR42
	m.RAM[0xFF22] = 0x00 //NR43
	m.RAM[0xFF23] = 0xBF //NR30
	m.RAM[0xFF24] = 0x77 //NR50
	m.RAM[0xFF25] = 0xF3 //NR51
	m.RAM[0xFF26] = 0xF1 //NR52
	m.RAM[0xFF40] = 0x91 //LCDC
	m.RAM[0xFF42] = 0x00 //SCY
	m.RAM[0xFF43] = 0x00 //SCX
	m.RAM[0xFF45] = 0x00 //LYC
	m.RAM[0xFF47] = 0xFC //BGP
	m.RAM[0xFF48] = 0xFF //OBP0
	m.RAM[0xFF49] = 0xFF //OBP1
	m.RAM[0xFF4A] = 0x00 //WY
	m.RAM[0xFF4B] = 0x00 //WX
	m.RAM[0xFFFF] = 0x00 //IE
}

func (m *Memory) Write(pointer Address, i byte) {
	m.RAM[pointer] = i
}

func (m *Memory) Read(addr Address) byte {
	return m.RAM[addr]
}
