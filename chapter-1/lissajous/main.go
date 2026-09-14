package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

const (
	whiteIndex = 0 // first color in palette
	greenIndex = 1 // next color in palette
	redIndex   = 2 // next color in palette
	blueIndex  = 3 // next color in palette
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0, 255, 0, 1}, // green
	color.RGBA{255, 0, 0, 1}, // red
	color.RGBA{0, 0, 255, 1}, // blue
}

/*
```sh
go run main.go > out.gif
```
*/
func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for range nframes {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(
				size+int(x*size+0.5),
				size+int(y*size+0.5),
				greenIndex)
			img.SetColorIndex(size+int(x*size+0.5),
				size+int(y*size+0.5),
				redIndex)
			img.SetColorIndex(size+int(x*size+0.9),
				size+int(y*size+0.9),
				blueIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
