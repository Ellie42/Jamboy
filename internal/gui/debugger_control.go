package gui

import (
	"git.agehadev.com/elliebelly/gooey/lib/eventmanager"
	"git.agehadev.com/elliebelly/jamboy/internal/engine"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type DebuggerControl struct {
	jamboy       *engine.Jamboy
	eventManager *eventmanager.EventManager
}

func NewDebuggerControl(jb *engine.Jamboy, e *eventmanager.EventManager) *DebuggerControl {
	e.AddOnKeyDownListener(glfw.KeyF8, func(key glfw.Key) {
		jb.Debugger.Step()
	})

	e.AddOnKeyDownListener(glfw.KeyF9, func(key glfw.Key) {
		jb.Debugger.Continue()
	})

	return &DebuggerControl{
		jamboy:       jb,
		eventManager: e,
	}
}
