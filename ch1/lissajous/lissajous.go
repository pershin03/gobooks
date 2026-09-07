package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.Black,                        // индекс 0 — фон
	color.White,                        // индекс 1 — белый
	color.RGBA{0x39, 0xFF, 0x14, 0xFF}, // индекс 2 — фосфорно-зелёный (неоновый)
	color.RGBA{0xFF, 0x00, 0x00, 0xFF}, // индекс 3 — красный
	color.RGBA{0xFF, 0x7F, 0x00, 0xFF}, // индекс 4 — оранжевый
	color.RGBA{0xFF, 0xFF, 0x00, 0xFF}, // индекс 5 — жёлтый
	color.RGBA{0x00, 0xFF, 0xFF, 0xFF}, // индекс 6 — голубой
	color.RGBA{0x00, 0x00, 0xFF, 0xFF}, // индекс 7 — синий
	color.RGBA{0xFF, 0x00, 0xFF, 0xFF}, // индекс 8 — фиолетовый
}

const (
	blackIndex    = 0
	whiteIndex    = 1
	neoGreenIndex = 2
)

func main() {
	f, err := os.Create("out.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lissajous(f)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	numColors := uint8(len(palette) - 1)
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		step := 0
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			colorIndex := uint8(step)%numColors + 1
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				colorIndex)
			step++
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
