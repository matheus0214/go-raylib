package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/matheus0214/go-raylib/configs"
	"github.com/matheus0214/go-raylib/environment"
	"github.com/matheus0214/go-raylib/player"
)

func main() {
	rl.InitWindow(configs.WINDOW_WIDTH, configs.WINDOW_HEIGHT, configs.GAME_NAME)
	defer rl.CloseWindow()

	rl.SetTargetFPS(configs.FPS)

	player.LoadPlayer()
	environment.Load()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.White)

		environment.Draw()
		player.DrawPlayer()

		rl.EndDrawing()
	}
}
