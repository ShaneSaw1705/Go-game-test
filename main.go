package main

import (
	"example.com/gogame/player"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(800, 450, "Game")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	plr := player.NewPlayer(20, 20, 5, 10)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		plr.Move()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("This is my game written in golang", 190, 200, 20, rl.LightGray)

		plr.Render()

		rl.EndDrawing()
	}
}
