package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	fmt.Println("Starting program...")
	rl.InitWindow(800, 600, "Sand Game")
	fmt.Println("Window initialized")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		//rl.DrawText("Ding Dong Darnell at your service", 10, 10, 20, rl.White)
		rl.EndDrawing()
	}

	fmt.Println("Closing...")
	rl.CloseWindow()
}
