package main

import (
	"os"
	"sync"

	"main/plugins"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/kbinani/screenshot"
)

var (
	TargetFPS int32   = 0
	movFactor         = rl.GetFrameTime() * speed
	speed     float32 = 300

	plugMap map[string]*plugin
)

type frame struct {
	x, y int
}

// ##[plugins.Plugin]
type plugin struct {
	Name string

	// false by default
	// you must to enable to run with: ##[loadPlugins]
	Enabled bool // false

	// redirect the function to your plugin
	//
	// u can make this with anonynmous function
	//
	// with opitional concurrency :)
	//
	// go func() { ... }
	Fn func()
}

func stdPlugins() {
	for _, p := range plugMap {
		if p.Enabled {
			p.Fn()
		}
	}
}

func pluginByName(name string) {
	for i, p := range plugMap {
		if name == i {
			p.Fn()
			return
		}
	}
}

func loadPlugins(name ...string) {
	if len(name) == 0 {
		stdPlugins()
		return
	}
	pluginByName(name[0])
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
			Name:  "DVD",
			Size:  .2,
			Path:  "../assets/DVD.png",
			Color: rl.White,
			Mu:    sync.Mutex{},
		},
	)
	// TODO: create wrapper func like Fn() in plugin's package
	// to control settings
	plugMap = map[string]*plugin{
		"Grid": {
			Enabled: false,
			Fn:      func() { plugins.Grid() },
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
		if rl.IsKeyPressed(rl.KeyM) {
			plugMap["Movement"].Enabled = !plugMap["Movement"].Enabled
		}
		if !plugMap["Color"].Enabled {
			image.Color = rl.White
		}
		loadPlugins()
		rl.DrawTexture(image.Texture, int32(image.PosX), int32(image.PosY), image.Color)
		// rl.DrawTexture(image2.Texture, int32(image2.PosX), int32(image2.PosY), image2.Color)

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
