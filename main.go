package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/matheus0214/go-raylib/configs"
)

func main() {
	rl.InitWindow(configs.WINDOW_WIDTH, configs.WINDOW_HEIGHT, configs.GAME_NAME)
	defer rl.CloseWindow()

	rl.SetTargetFPS(configs.FPS)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.White)
		rl.DrawText("Congrats", 190, 100, 20, rl.Black)

		rl.EndDrawing()
	}
}
