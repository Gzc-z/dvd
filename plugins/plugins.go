// Package plugins
package plugins

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Plugin struct {
	Name   string
	Enable bool
	Fn     func()
}

// Image is a configuration for NewImage()
// TODO: add comments
type Image struct {
	PosX, PosY float32
	W, H       int32

	Path string

	Speed float32

	// check which face is pointed
	faceR, faceD bool

	// Resize image to fit the given size
	// default value is 1
	Size float64

	Texture rl.Texture2D
}

func LoadImage(i *Image) *Image {
	image := rl.LoadImage(i.Path)

	if i.Size == 0 {
		i.Size = 1
	}

	var w, h int32

	if i.W != 0 {
		w = int32(float64(i.W) * i.Size)
	} else {
		w = int32(float64(image.Width) * i.Size)
	}

	if i.H != 0 {
		h = int32(float64(i.H) * i.Size)
	} else {
		h = int32(float64(image.Height) * i.Size)
	}

	rl.ImageResize(image, w, h)
	i.Texture = rl.LoadTextureFromImage(image)
	rl.UnloadImage(image)
	return i
}

func Movement(i *Image) {
	moveX(i)
	moveY(i)
}

func Grid() {
	ggap := 20
	gap := 10
	for i := range rl.GetScreenHeight()/ggap + gap {
		rl.DrawLine(int32(gap*-1), int32(i*ggap), int32(rl.GetScreenWidth()), int32(i*ggap), rl.Black)
	}
	for i := range rl.GetScreenWidth()/ggap + gap {
		rl.DrawLine(int32(i*ggap), int32(gap*-1), int32(i*ggap), int32(rl.GetScreenHeight()), rl.Black)
	}
}

func moveY(i *Image) {
	if i.PosY > float32(rl.GetScreenHeight())-float32(i.Texture.Height) {
		i.faceD = !i.faceD
	} else if i.PosY < 0 && i.faceD {
		i.faceD = !i.faceD
	}
	if !i.faceD {
		i.PosY += i.Speed
		return
	}
	i.PosY -= i.Speed
}

func moveX(i *Image) {
	if i.PosX > float32(rl.GetScreenWidth())-float32(i.Texture.Width) {
		i.faceR = !i.faceR
	} else if i.PosX < 0 && i.faceR {
		i.faceR = !i.faceR
	}
	if !i.faceR {
		i.PosX += i.Speed
		return
	}
	i.PosX -= i.Speed
}

func randomColor() int {
	return rand.Intn(0xffffff)
}
