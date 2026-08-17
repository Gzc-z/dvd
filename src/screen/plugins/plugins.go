// Package plugins
package plugins

import (
	"image/color"
	"math/rand"
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Plugin struct {
	Name string

	// its false by default
	Enable bool

	// redirect the function to your plugin
	//
	// u can make this with anonynmous function
	// func() { ... }
	Fn func()
}

// Image is a configuration for NewImage()
// TODO: add comments
type Image struct {
	Name       string
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
	Mu    sync.Mutex
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

var (
	// ##[Grid]
	gridSize = 30

	// dont use nevative values
	offboard = 20
)

// Grid draws a grid on the screen
func Grid(customOffboard ...int) {
	if customOffboard != nil {
		offboard = customOffboard[0]
	}

	height := rl.GetScreenHeight()
	width := rl.GetScreenWidth()

	// horizontal lines
	for i := range height/gridSize + offboard {
		rl.DrawLine(
			int32(-offboard),      // <- (startPosX)
			int32(i*gridSize),     // <- (startPosY)
			int32(width+offboard), // <- (endPosX)
			int32(i*gridSize),     // <- (endPosY)
			rl.White,
		)
	}
	// vertical lines
	for i := range width/gridSize + offboard {
		rl.DrawLine(
			int32(i*gridSize),      // <- (startPosX)
			int32(-offboard),       // <- (startPosY)
			int32(i*gridSize),      // <- (endPosX)
			int32(height+offboard), // <- (endPosY)
			rl.White,
		)
	}
}

var timer float64

func Color(i *Image) {
	timer += float64(rl.GetFrameTime())
	if timer >= .1 {
		timer = 0
		i.Mu.Lock()
		i.Color = randomColor()
		i.Mu.Unlock()
	}
}

func Movement(i *Image) {
	moveX(i)
	moveY(i)
}

func moveY(i *Image) {
	hLimit := float32(int32(rl.GetScreenHeight()) - i.Texture.Height)
	if i.PosY > hLimit && !i.faceD {
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
	wLimit := float32(int32(rl.GetScreenWidth()) - i.Texture.Width)
	if i.PosX > wLimit && !i.faceR {
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
		A: 0xff,
	}
	// fmt.Printf("%#v\n", color)
	return color
}
