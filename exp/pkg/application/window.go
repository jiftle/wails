package application

import (
	"sync/atomic"

	"github.com/wailsapp/wails/exp/pkg/options"
)

type windowImpl interface {
	setTitle(title string)
	setSize(width, height int)
	setAlwaysOnTop(alwaysOnTop bool)
	run() error
	navigateToURL(url string)
	setResizable(resizable bool)
	setMinSize(width, height int)
	setMaxSize(width, height int)
	enableDevTools()
	execJS(js string)
	setMaximised()
	setMinimised()
	setFullscreen()
	isMinimised() bool
	isMaximised() bool
	isFullscreen() bool
	restore()
	setBackgroundColor(color *options.RGBA)
}

type Window struct {
	options *options.Window
	impl    windowImpl
	id      uint64
}

var windowID atomic.Uint64

func NewWindow(options *options.Window) *Window {
	id := windowID.Load()
	windowID.Add(1)
	return &Window{
		id:      id,
		options: options,
	}
}

func (w *Window) SetTitle(title string) {
	if w.impl == nil {
		w.options.Title = title
		return
	}
	w.impl.setTitle(title)
}

func (w *Window) SetSize(width, height int) {
	if w.impl == nil {
		w.options.Width = width
		w.options.Height = height
		return
	}
	w.impl.setSize(width, height)
}

func (w *Window) Run() error {
	w.impl = newWindowImpl(w.options)
	return w.impl.run()
}

func (w *Window) SetAlwaysOnTop(b bool) {
	if w.impl == nil {
		w.options.AlwaysOnTop = b
		return
	}
	w.impl.setAlwaysOnTop(b)
}

func (w *Window) NavigateToURL(s string) {
	if w.impl == nil {
		w.options.URL = s
		return
	}
	w.impl.navigateToURL(s)
}

func (w *Window) SetResizable(b bool) {
	if w.impl == nil {
		w.options.DisableResize = !b
		return
	}
	w.impl.setResizable(b)
}

func (w *Window) SetMinSize(minWidth, minHeight int) {
	if w.impl == nil {
		w.options.MinWidth = minWidth
		if w.options.Width < minWidth {
			w.options.Width = minWidth
		}
		w.options.MinHeight = minHeight
		if w.options.Height < minHeight {
			w.options.Height = minHeight
		}
		return
	}
	w.impl.setSize(w.options.Width, w.options.Height)
	w.impl.setMinSize(minWidth, minHeight)
}
func (w *Window) SetMaxSize(maxWidth, maxHeight int) {
	if w.impl == nil {
		w.options.MinWidth = maxWidth
		if w.options.Width > maxWidth {
			w.options.Width = maxWidth
		}
		w.options.MinHeight = maxHeight
		if w.options.Height > maxHeight {
			w.options.Height = maxHeight
		}
		return
	}
	w.impl.setSize(w.options.Width, w.options.Height)
	w.impl.setMaxSize(maxWidth, maxHeight)
}

func (w *Window) EnableDevTools() {
	if w.impl == nil {
		w.options.EnableDevTools = true
		return
	}
	w.impl.enableDevTools()
}

func (w *Window) ExecJS(js string) {
	if w.impl == nil {
		return
	}
	w.impl.execJS(js)
}

// Set Maximized
func (w *Window) SetMaximized() {
	if w.impl == nil {
		w.options.StartState = options.WindowStateMaximised
		return
	}
	w.impl.setMaximised()
}

// Set Minimized
func (w *Window) SetMinimized() {
	if w.impl == nil {
		w.options.StartState = options.WindowStateMinimised
		return
	}
	w.impl.setMinimised()
}

// Set Fullscreen
func (w *Window) SetFullscreen() {
	if w.impl == nil {
		w.options.StartState = options.WindowStateFullscreen
		return
	}
	w.impl.setFullscreen()
}

// IsMinimised returns true if the window is minimised
func (w *Window) IsMinimised() bool {
	if w.impl == nil {
		return false
	}
	return w.impl.isMinimised()
}

// IsMaximised returns true if the window is maximised
func (w *Window) IsMaximised() bool {
	if w.impl == nil {
		return false
	}
	return w.impl.isMaximised()
}

// IsFullscreen returns true if the window is fullscreen
func (w *Window) IsFullscreen() bool {
	if w.impl == nil {
		return false
	}
	return w.impl.isFullscreen()
}

func (w *Window) SetBackgroundColor(color *options.RGBA) {
	if w.impl == nil {
		w.options.BackgroundColour = color
		return
	}
	w.impl.setBackgroundColor(color)
}
