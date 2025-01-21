package sprites

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"errors"
	"fmt"
)

func LoadSprite(path string) (rl.Texture2D, error) {
	texture := rl.LoadTexture(path)
	if texture.ID == 0 {
		return rl.NewTexture2D(0, 0, 0, 0, 0), errors.New(fmt.Sprintf("\033[0;31mError to load sprite: %s\033[0m", path))
	}

	return texture, nil
}

func UnloadTexture(texture rl.Texture2D) {
	rl.UnloadTexture(texture)
}

