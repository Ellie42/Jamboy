package engine

import "fmt"

type GPU struct {
	Clocks             uint
	currentStateClocks int
	currentState       GPUState
	currentRow         int
	jb                 *Jamboy
	lcdControlFlags    byte
	PixelBuffer        *[]uint8
	backgroundMaps     [][]byte
	tileDataBlocks     [][]byte
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
	tileMap := g.backgroundMaps[0]

	if g.jb.MMU.RAM[AddrLCDControl]&byte(LCDCBGTileMapDisplay) > 0 {
		tileMap = g.backgroundMaps[1]
	}

	tileData := g.tileDataBlocks[0]

	tileMapLow := g.jb.MMU.RAM[AddrLCDControl]&byte(LCDCBGTileMapDisplay) > 0

	if tileMapLow {
		tileData = g.tileDataBlocks[1]
	}

	currentRow := g.currentRow

	for x := 0; x < ResolutionX; x++ {
		xTileNum := x / 8
		yTileNum := ((ResolutionY) - currentRow) / 8
		xRemainder := x - xTileNum*8
		yRemainder := ((ResolutionY) - currentRow - yTileNum*8)
		currentTile := xTileNum + (yTileNum * 32)
		tilePosition := tileMap[currentTile]

		if tilePosition > 0 {
			fmt.Println(tilePosition)
		}

		var tileByteOffset int

		if tileMapLow {
			tileByteOffset = int(tilePosition*16 + byte(yRemainder)*2)
		} else {
			tileByteOffset = int(int8(tilePosition))*16 + int(uint8(yRemainder))*2
		}

		bitShiftAmount := uint8(7 - xRemainder)
		bitMask := byte(1 << bitShiftAmount)

		blow := (tileData[tileByteOffset] & bitMask) >> bitShiftAmount
		bhigh := (tileData[tileByteOffset+1] & bitMask) >> (bitShiftAmount)

		palette := blow | (bhigh << 1)

		currentI := x*4 + (currentRow * ResolutionX * 4)

		if palette > 0 {
			(*g.PixelBuffer)[currentI] = 0x00
			(*g.PixelBuffer)[currentI+1] = 0x00
			(*g.PixelBuffer)[currentI+2] = 0x00
			(*g.PixelBuffer)[currentI+3] = 0xFF
		} else {
			(*g.PixelBuffer)[currentI] = 0xFF
			(*g.PixelBuffer)[currentI+1] = 0xFF
			(*g.PixelBuffer)[currentI+2] = 0xFF
			(*g.PixelBuffer)[currentI+3] = 0xFF
		}
	}
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
		backgroundMaps: [][]byte{
			jamboy.MMU.RAM[BackgroundMap0.From : BackgroundMap0.To+1],
			jamboy.MMU.RAM[BackgroundMap1.From : BackgroundMap1.To+1],
		},
		tileDataBlocks: [][]byte{
			jamboy.MMU.RAM[TileBlock0.From : TileBlock1.To+1],
			jamboy.MMU.RAM[TileBlock1.From : TileBlock2.To+1],
		},
	}
}
