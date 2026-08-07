package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/kbinani/screenshot"
)

var (
	once      sync.Once
	texture   rl.Texture2D
	mov       float32
	TargetFPS int32   = 75
	speed     float32 = 200
)

type frame struct {
	x, y int
}

// Image is a configuration for NewImage()
type Image struct {
	PosX, PosY float32
	W, H       int32
	Path       string

	speed float32

	// check which face is pointed
	FaceR, FaceD bool

	// Resize image to fit the given size
	// default value is 1
	Size float64
}

func randomColor() int {
	return rand.Intn(0xffffff)
}

func grid() {
	ggap := 20
	gap := 10
	for i := range rl.GetScreenHeight()/ggap + gap {
		rl.DrawLine(int32(gap*-1), int32(i*ggap), int32(rl.GetScreenWidth()), int32(i*ggap), rl.Black)
	}
	for i := range rl.GetScreenWidth()/ggap + gap {
		rl.DrawLine(int32(i*ggap), int32(gap*-1), int32(i*ggap), int32(rl.GetScreenHeight()), rl.Black)
	}
}

func mvy(i *Image) {
	if i.PosY > float32(rl.GetScreenHeight())-float32(texture.Height) {
		i.FaceD = !i.FaceD
	} else if i.PosY < 0 && i.FaceD {
		i.FaceD = !i.FaceD
	}
	if !i.FaceD {
		i.PosY += mov
		return
	}
	i.PosY -= mov
}

func mvx(i *Image) {
	if i.PosX > float32(rl.GetScreenWidth())-float32(texture.Width) {
		i.FaceR = !i.FaceR
	} else if i.PosX < 0 && i.FaceR {
		i.FaceR = !i.FaceR
	}
	if !i.FaceR {
		i.PosX += mov
		return
	}
	i.PosX -= mov
}

func LoadImage(i *Image) *Image {
	image := rl.LoadImage(i.Path)

	if i.Size == 0 {
		i.Size = 1
	}

	var w, h int32
	if i.W != 0 && i.H != 0 {
		w = int32(float64(i.W) * i.Size)
		h = int32(float64(i.H) * i.Size)
	} else {
		w = int32(float64(image.Width) * i.Size)
		h = int32(float64(image.Height) * i.Size)
	}
	rl.ImageResize(image, w, h)

	texture = rl.LoadTextureFromImage(image)
	rl.UnloadImage(image)
	return i
}

func settings() {
	once.Do(func() {
		mov = speed * rl.GetFrameTime()
	})
	fmt.Println(mov)
	os.Exit(0)
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

	image := LoadImage(
		&Image{
			Size: .2,
			Path: "assets/DVD.png",
		},
	)
	defer rl.UnloadTexture(texture)

	rl.SetTargetFPS(TargetFPS)
	rl.SetExitKey(rl.KeyQ)
	for !rl.WindowShouldClose() {
		go settings()
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawFPS(0, 0)

		go mvx(image)
		go mvy(image)
		// grid()
		// fmt.Println(mov)
		rl.DrawTexture(texture, int32(image.PosX), int32(image.PosY), rl.White)

		rl.EndDrawing()
	}
	ctx <- os.Interrupt
	close(ctx)
}

func main() {
	ctx := make(chan os.Signal, 1)

	initWindow(ctx)
}
