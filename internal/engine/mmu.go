package engine

type MemoryRange struct {
	From Address
	To   Address
}

var (
	CartROM0 = MemoryRange{
		From: Address(0x0000),
		To:   Address(0x3FFF),
	}
	CartROMN = MemoryRange{
		From: Address(0x3FFF),
		To:   Address(0x7FFF),
	}
	VRAM = MemoryRange{
		From: 0x8000,
		To:   0x9FFF,
	}
	TileData = MemoryRange{
		From: Address(0x8000),
		To:   Address(0x97FF),
	}
	TileBlock0 = MemoryRange{
		From: Address(0x8000),
		To:   Address(0x87FF),
	}
	TileBlock1 = MemoryRange{
		From: Address(0x8800),
		To:   Address(0x8FFF),
	}
	TileBlock2 = MemoryRange{
		From: Address(0x9000),
		To:   Address(0x97FF),
	}
	BackgroundMap0 = MemoryRange{
		From: Address(0x9800),
		To:   Address(0x9BFF),
	}
	BackgroundMap1 = MemoryRange{
		From: Address(0x9C00),
		To:   Address(0x9FFF),
	}
)

const (
	AddrLCDControl     Address = 0xFF40
	AddrLCDCStatus             = 0xFF41
	AddrScrollY                = 0xFF42
	AddrScrollX                = 0xFF43
	AddrLY                     = 0xFF44
	AddrLYC                    = 0xFF45
	AddrBootROMDisable         = 0xFF50
)

type Address uint16

func (a Address) InRange(r MemoryRange) bool {
	if a >= r.From && a <= r.To {
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
type MMU struct {
	Jamboy *Jamboy
	RAM    [0x10000]byte
}

func (m *MMU) Reset() {
	for i := 0; i < 0xFFFF; i++ {
		m.RAM[i] = 0
	}

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

func (m *MMU) Write(addr Address, i byte) {
	m.Jamboy.CPU.Wait(1)

	if addr.InRange(VRAM) {
		//runtime.Breakpoint()
	}

	m.RAM[addr] = i
}

func (m *MMU) Read(addr Address) byte {
	m.Jamboy.CPU.Wait(1)

	return m.ReadInstant(addr)
}

func (m *MMU) ReadInstant(addr Address) byte {
	if m.Jamboy.CPU.IsBooting && addr < 0x0100 {
		return m.RAM[addr]
	}

	if (addr.InRange(CartROM0) || addr.InRange(CartROMN)) && int(addr) < len(m.Jamboy.Cart.Data) {
		return m.Jamboy.Cart.Data[addr]
	}

	return m.RAM[addr]
}

func newMMU(jb *Jamboy) *MMU {
	return &MMU{
		Jamboy: jb,
		RAM:    [65536]byte{},
	}
}
