// Package plugins
package plugins

import (
	"image/color"
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
	Image   *rl.Image

	Color color.RGBA
}

func (i *Image) LoadImage() {
	i.Image = rl.LoadImageFromTexture(i.Texture)
}

func (i *Image) LoadTexture() {
	i.Texture = rl.LoadTextureFromImage(i.Image)
}

func GetImage(i *Image) *Image {
	i.Image = rl.LoadImage(i.Path)

	if i.Size == 0 {
		i.Size = 1
	}

	var w, h int32

	if i.W != 0 {
		w = int32(float64(i.W) * i.Size)
	} else {
		w = int32(float64(i.Image.Width) * i.Size)
	}

	if i.H != 0 {
		h = int32(float64(i.H) * i.Size)
	} else {
		h = int32(float64(i.Image.Height) * i.Size)
	}

	rl.ImageResize(i.Image, w, h)
	i.Texture = rl.LoadTextureFromImage(i.Image)
	rl.UnloadImage(i.Image)
	return i
}

func Grid() {
	ggap := 20
	gap := 10
	for i := range rl.GetScreenHeight()/ggap + gap {
		rl.DrawLine(int32(gap*-1), int32(i*ggap), int32(rl.GetScreenWidth()), int32(i*ggap), rl.White)
	}
	for i := range rl.GetScreenWidth()/ggap + gap {
		rl.DrawLine(int32(i*ggap), int32(gap*-1), int32(i*ggap), int32(rl.GetScreenHeight()), rl.White)
	}
}

func Color(i *Image) {
	i.Color = randomColor()
}

func Movement(i *Image) {
	moveX(i)
	moveY(i)
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

func randomColor() color.RGBA {
	color := color.RGBA{
		R: uint8(rand.Intn(0xFF)),
		G: uint8(rand.Intn(0xFF)),
		B: uint8(rand.Intn(0xFF)),
		A: uint8(rand.Intn(0xFF)),
	}
	// fmt.Printf("%#v\n", color)
	return color
}
