package main

import (
	"os"

	"main/plugins"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/kbinani/screenshot"
)

var (
	TargetFPS int32   = 256
	movFactor         = rl.GetFrameTime() * speed
	speed     float32 = 300

	plugMap map[string]*plugin
)

type frame struct {
	x, y int
}

type plugin struct {
	Enabled bool
	Fn      func()
}

func loadPlugins() {
	for _, p := range plugMap {
		if p.Enabled {
			p.Fn()
		}
	}
}

func initWindow(ctx chan<- os.Signal) {
	resolution := screenshot.GetDisplayBounds(0)
	frame := frame{
		x: resolution.Dx() * 6 / 10,
		y: resolution.Dy() * 6 / 10,
	}

	rl.InitWindow(
		int32(frame.x),
		int32(frame.y),
		"bouncing DVD",
	)
	defer rl.CloseWindow()
	image := plugins.GetImage(
		&plugins.Image{
			Size:  .2,
			Path:  "assets/DVD.png",
			Color: rl.White,
		},
	)
	plugMap = map[string]*plugin{
		"Grid": {
			Enabled: false,
			Fn:      plugins.Grid,
		},
		"Movement": {
			Enabled: true,
			Fn:      func() { plugins.Movement(image) },
		},
		"Color": {
			Enabled: false,
			Fn:      func() { plugins.Color(image) },
		},
	}

	rl.SetTargetFPS(TargetFPS)
	rl.SetExitKey(rl.KeyQ)
	for !rl.WindowShouldClose() {
		image.Speed = speed * rl.GetFrameTime()
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawFPS(0, 0)

		if rl.IsKeyPressed(rl.KeyG) {
			plugMap["Grid"].Enabled = !plugMap["Grid"].Enabled
		}
		if rl.IsKeyPressed(rl.KeyC) {
			plugMap["Color"].Enabled = !plugMap["Color"].Enabled
		}
		if !plugMap["Color"].Enabled {
			image.Color = rl.White
		}
		loadPlugins()
		rl.DrawTexture(image.Texture, int32(image.PosX), int32(image.PosY), image.Color)

		rl.EndDrawing()
	}
	rl.UnloadTexture(image.Texture)
	ctx <- os.Interrupt
	close(ctx)
}

func main() {
	ctx := make(chan os.Signal, 1)

	initWindow(ctx)
}
