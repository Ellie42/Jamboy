package engine

type GPU struct {
	Clocks             uint
	currentStateClocks int
	currentState       GPUState
	currentRow         int
	jb                 *Jamboy
	lcdControlFlags    byte
}

type LCDControlFlag byte

const (
	LCDCBGDisplayEnable LCDControlFlag = 1 << iota
	LCDCOBJDisplayEnable
	LCDCOBJSize
	LCDCBGTileMapDisplay
	LCDCBGWindowTileDataSelect
	LCDCWindowDisplayEnable
	LCDCWindowTileMapDisplayData
	LCDCDisplayEnable
)

const (
	ResolutionX = 160
	ResolutionY = 144
)

type GPUState uint

const (
	OAM GPUState = iota
	OAMVRAM
	HBlank
	VBlank
)

var clockCounts = map[GPUState]uint{
	HBlank:  201,
	OAM:     77,
	OAMVRAM: 169,
	VBlank:  456,
}

var stateFlags = []byte{
	0,
	2,
	3,
	1,
}

func (g *GPU) Tick() {
	g.lcdControlFlags = g.jb.MMU.ReadInstant(AddrLCDControl)

	if !g.LCDCFlagEnabled(LCDCDisplayEnable) {
		g.Clocks = 0
		return
	}

	g.currentStateClocks -= int(g.Clocks)
	g.Clocks = 0

	for g.currentStateClocks < 0 {
		if g.currentState == VBlank {
			if g.currentRow < 153 {
				g.IncrementRow()
			} else {
				g.currentState = OAM
				g.ResetRow()
			}
		} else {
			g.currentState = GPUState((uint(g.currentState) + 1) % 3)
		}

		if g.currentState == HBlank {
			g.DrawRow()
		} else if g.currentState == OAM {
			if g.currentRow == 143 {
				g.currentState = VBlank
				//fmt.Printf("Vblank!\n")
			}

			g.IncrementRow()
		}

		//Setting GPU mode flag
		g.jb.MMU.RAM[AddrLCDCStatus] = (g.jb.MMU.RAM[AddrLCDCStatus] & (^uint8(0x03))) | stateFlags[g.currentState]

		g.currentStateClocks += int(clockCounts[g.currentState])
	}
}

func (g *GPU) DrawRow() {
	//fmt.Printf("Drawing row\n")
}

func (g *GPU) IncrementRow() {
	g.currentRow++
	g.writeRowValues()
}

func (g *GPU) ResetRow() {
	g.currentRow = 0
	g.writeRowValues()
}

func (g *GPU) Reset() {
	g.ResetRow()
}

func (g *GPU) writeRowValues() {
	g.jb.MMU.Write(AddrLY, byte(g.currentRow))

	coincidenceFlag := uint8(0)

	if g.currentRow == int(g.jb.MMU.ReadInstant(AddrLYC)) {
		coincidenceFlag = 0x04

		LYCInterrupt()
	}

	// Setting LY == LYC coincidence flag
	g.jb.MMU.RAM[AddrLCDCStatus] = (g.jb.MMU.RAM[AddrLCDCStatus] & (^uint8(0x04))) | coincidenceFlag
}

func (g GPU) LCDCFlagEnabled(flag LCDControlFlag) bool {
	return g.jb.MMU.ReadInstant(AddrLCDControl)&byte(flag) > 0
}

func LYCInterrupt() {
	//fmt.Println("LYC Interrupt!")
}

func NewGPU(jamboy *Jamboy) *GPU {
	return &GPU{
		currentStateClocks: int(clockCounts[OAM]),
		jb:                 jamboy,
	}
}
