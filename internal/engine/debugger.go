package engine

import "sync"

type Debugger struct {
	breakPoints map[int]bool

	jamboy   *Jamboy
	cont     bool
	haltNext bool
	sync.Mutex
}

func (d *Debugger) Reset() {

}

func (d *Debugger) PreTick() {
	d.Lock()
	defer d.Unlock()

	if d.cont {
		d.cont = false
		d.jamboy.IsHalted = false
	} else if _, ok := d.breakPoints[int(d.jamboy.CPU.PC-1)]; ok {
		d.jamboy.IsHalted = true
	} else if d.haltNext {
		d.jamboy.IsHalted = true
		d.haltNext = false
	}
}

func (d *Debugger) PostTick() {
}

func (d *Debugger) Step() {
	d.Lock()
	defer d.Unlock()

	d.haltNext = true
	d.cont = true
}

func (d *Debugger) Continue() {
	d.Lock()
	defer d.Unlock()

	d.cont = true
}

func (d *Debugger) AddBreakpoint(address int) {
	d.Lock()
	defer d.Unlock()

	d.breakPoints[address] = true
}

func (d *Debugger) RemoveBreakpoint(address int) {
	d.Lock()
	defer d.Unlock()

	delete(d.breakPoints, address)
}

func (d *Debugger) GetCurrentLine() int {
	return d.jamboy.Code.GetLineForAddress(int(d.jamboy.CPU.PC - 1))
}

func NewDebugger(jb *Jamboy) *Debugger {
	return &Debugger{
		jamboy:      jb,
		breakPoints: make(map[int]bool),
	}
}
