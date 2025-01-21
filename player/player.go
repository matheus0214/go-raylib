package player

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/matheus0214/go-raylib/sprites"
	"log"
)

const (
	dirDown  = 0
	dirUp    = 16
	dirLeft  = 32
	dirRight = 48

	speed       = 3
	framesSpeed = 3
)

var (
	playerSprite     rl.Texture2D
	playerSpriteSrc  rl.Rectangle
	playerSpriteDest rl.Rectangle

	direction = "down"

	framesCounter   = 0
	playerFrame     = 0
	playerDirection = 0
	isMoving        = false
)

func LoadPlayer() {
	var err error
	playerSprite, err = sprites.LoadSprite("res/character/boy.png")
	if err != nil {
		log.Fatal(err)
	}

	playerSpriteSrc = rl.NewRectangle(0, 0, 16, 16)
	playerSpriteDest = rl.NewRectangle(0, 0, 48, 48)
}

func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		direction = "up"
		isMoving = true
	}

	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		direction = "down"
		isMoving = true
	}

	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		direction = "left"
		isMoving = true
	}

	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		direction = "right"
		isMoving = true
	}
}

func movement() {
	if isMoving {
		framesCounter++

		switch direction {
		case "up":
			playerSpriteDest.Y -= speed
			playerDirection = dirUp
		case "down":
			playerSpriteDest.Y += speed
			playerDirection = dirDown
		case "left":
			playerSpriteDest.X -= speed
			playerDirection = dirLeft
		case "right":
			playerSpriteDest.X += speed
			playerDirection = dirRight
		}

		if playerFrame > 3 {
			playerFrame = 0
		}

		if framesCounter%10 == 1 {
			playerFrame++
		}
	} else {
		playerFrame = 0
	}

	isMoving = false
}

func DrawPlayer() {
	input()
	movement()

	playerSpriteSrc.X = float32(playerDirection)
	playerSpriteSrc.Y = playerSpriteSrc.Height * float32(playerFrame)

	rl.DrawTexturePro(playerSprite, playerSpriteSrc, playerSpriteDest, rl.NewVector2(0, 0), 0, rl.White)
}