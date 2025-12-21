package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/betasve/programming-presentation/input"
	"github.com/betasve/programming-presentation/screen"
)

const (
	windowWidth  = 1280
	windowHeight = 720
	windowTitle  = "Presentation"
)

func main() {
	rl.InitWindow(windowWidth, windowHeight, windowTitle)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	manager := screen.NewManager()

	// Add sample screens for testing
	manager.Add(&BlankScreen{color: rl.DarkGray, label: "Slide 1 - Press Right Arrow or Space"})
	manager.Add(&BlankScreen{color: rl.DarkBlue, label: "Slide 2"})
	manager.Add(&BlankScreen{color: rl.DarkGreen, label: "Slide 3"})

	manager.Start()

	for !rl.WindowShouldClose() {
		input.Handle(manager)
		manager.Update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		manager.Draw()

		// Draw slide counter
		counter := fmt.Sprintf("%d / %d", manager.CurrentIndex()+1, manager.Count())
		rl.DrawText(counter, windowWidth-100, windowHeight-40, 20, rl.White)

		rl.EndDrawing()
	}
}

// BlankScreen is a simple test screen with a solid color background.
type BlankScreen struct {
	color rl.Color
	label string
}

func (s *BlankScreen) Load()   {}
func (s *BlankScreen) Unload() {}

func (s *BlankScreen) Update() screen.Screen {
	return nil
}

func (s *BlankScreen) Draw() {
	rl.ClearBackground(s.color)

	// Draw centered label
	textWidth := rl.MeasureText(s.label, 40)
	x := (windowWidth - textWidth) / 2
	y := windowHeight / 2
	rl.DrawText(s.label, x, int32(y), 40, rl.White)

	// Draw navigation hints
	hints := "Arrow Keys: Navigate | F: Fullscreen | ESC: Exit"
	hintsWidth := rl.MeasureText(hints, 20)
	rl.DrawText(hints, (windowWidth-hintsWidth)/2, windowHeight-60, 20, rl.LightGray)
}
