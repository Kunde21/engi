package main

import (
	"github.com/paked/engi"
)

var (
	World *Game
)

type Game struct{}

func (game *Game) Preload() {
	engi.Files.Add("assets/326488.wav")
}

func (game *Game) Setup(w *engi.World) {
	engi.SetBg(0xFFFFFF)

	w.AddSystem(&engi.RenderSystem{})
	w.AddSystem(&engi.AudioSystem{})

	backgroundMusic := engi.NewEntity([]string{"AudioSystem"})
	backgroundMusic.AddComponent(&engi.AudioComponent{File: "326488.wav", Repeat: true, Background: true})

	w.AddEntity(backgroundMusic)
}

func main() {
	World = &Game{}
	engi.Open("Audio Demo", 1024, 640, false, World)
}
