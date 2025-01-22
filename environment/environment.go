package environment

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/matheus0214/go-raylib/sprites"
	"log"
)

var (
	tilesetField     rl.Texture2D
	tilesetFieldSrc  rl.Rectangle
	tilesetFieldDest rl.Rectangle
)

func Load() {
	var err error
	tilesetField, err = sprites.LoadSprite("res/maps/TilesetField.png")
	if err != nil {
		log.Fatal(err)
	}

	tilesetFieldSrc = rl.NewRectangle(0, 0, 16, 16)
	tilesetFieldDest = rl.NewRectangle(100, 250, 48, 48)
}

func Draw() {
	mapW := 10
	mapH := 5

	for h := range mapH {
		switch {
		case h == 0:
			tilesetFieldSrc.Y = float32(0)
		case h == mapH-1:
			tilesetFieldSrc.Y = float32(32)
		default:
			tilesetFieldSrc.Y = float32(h + 1*16)
		}

		for w := range mapW {
			switch {
			case w == 0:
				tilesetFieldSrc.X = float32(0)
			case w == mapW-1:
				tilesetFieldSrc.X = float32(32)
			default:
				tilesetFieldSrc.X = float32(w + 1*16)
			}

			tilesetFieldDest.X = 100 + float32(w*3)*tilesetFieldSrc.Width
			tilesetFieldDest.Y = 100 + float32(h*3)*tilesetFieldSrc.Height

			rl.DrawTexturePro(
				tilesetField, tilesetFieldSrc, tilesetFieldDest,
				rl.NewVector2(0, float32(0)),
				0, rl.White)
		}
	}
}
