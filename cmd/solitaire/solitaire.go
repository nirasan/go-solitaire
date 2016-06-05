package main

import (
	"flag"
	keyboard "github.com/jteeuwen/keyboard/termbox"
	"github.com/nirasan/solitaire/klondike"
	"github.com/nirasan/solitaire/klondike/renderer"
	"github.com/nsf/termbox-go"
)

type Renderer interface {
	Render()
	SetError(error)
}

var (
	k              *klondike.Klondike
	colorFlag      = flag.Bool("color", true, "draw color charactor")
	debugFlag      = flag.Bool("debug", false, "show debug")
	rendererString = flag.String("renderer", "basic", "termbox renderer")
	r              Renderer
)

func main() {
	flag.Parse()

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetOutputMode(termbox.Output256)

	k = klondike.NewKlondike()
	k.Init()

	switch *rendererString {
	case "basic":
		r = renderer.NewBasicRenderer(k, *colorFlag, *debugFlag)
	case "ls":
		r = renderer.NewLsRenderer(k, *colorFlag, *debugFlag)
	case "simple":
		r = renderer.NewSimpleRenderer(k, *colorFlag)
	}

	draw()

	pollEvent()
}

func draw() {
	r.Render()
}

func pollEvent() {
	running := true

	kb := keyboard.New()
	kb.Bind(func() { running = false }, "escape", "q")
	kb.Bind(func() { k.CursorUp(); draw() }, "up", "k")
	kb.Bind(func() { k.CursorDown(); draw() }, "down", "j")
	kb.Bind(func() { k.CursorLeft(); draw() }, "left", "h")
	kb.Bind(func() { k.CursorRight(); draw() }, "right", "l")
	kb.Bind(func() { k.CursorJump(); draw() }, "tab")
	kb.Bind(func() { r.SetError(k.Select()); draw() }, "space")

	for running {
		kb.Poll(termbox.PollEvent())
	}
}
