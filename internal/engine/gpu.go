package engine

import "fmt"

type GPU struct {
	Clocks             uint
	currentStateClocks int
	currentState       GPUState
	currentRow         int
	jb                 *Jamboy
}

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
	VBlank:  4560,
}

func (g *GPU) Tick() {
	g.currentStateClocks -= int(g.Clocks)
	g.Clocks = 0

	for g.currentStateClocks < 0 {
		if g.currentState == VBlank {
			g.currentState = OAM
			g.currentRow = 0
		} else {
			g.currentState = GPUState((uint(g.currentState) + 1) % 3)
		}

		if g.currentState == HBlank {
			g.DrawRow()
		} else if g.currentState == OAM {
			if g.currentRow == 143 {
				g.currentState = VBlank
				fmt.Printf("Vblank!\n")
			}

			g.currentRow++
		}

		g.currentStateClocks += int(clockCounts[g.currentState])
	}
}

func (g *GPU) DrawRow() {
	//fmt.Printf("Drawing row\n")
}

func NewGPU(jamboy *Jamboy) *GPU {
	return &GPU{
		currentStateClocks: int(clockCounts[OAM]),
		jb:                 jamboy,
	}
}
