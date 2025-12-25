package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/betasve/programming-presentation/input"
	"github.com/betasve/programming-presentation/screen"
)

const (
	windowWidth  = 1910
	windowHeight = 1020
	windowTitle  = "Presentation"
)

func main() {
	rl.InitWindow(windowWidth, windowHeight, windowTitle)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	manager := screen.NewManager()

	// Add screens
	manager.Add(&InitialScreen{backgroundPath: "assets/bg.png"})
	manager.Add(&ProgrammingScreen{backgroundPath: "assets/programming.png"})
	manager.Add(&BlankScreen{color: rl.DarkBlue, label: "Slide 2"})
	manager.Add(&BlankScreen{color: rl.DarkGreen, label: "Slide 3"})

	manager.SetTime(rl.GetTime())
	manager.Start()

	for !rl.WindowShouldClose() {
		manager.SetTime(rl.GetTime())
		input.Handle(manager)
		manager.Update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		manager.Draw()

		// Draw overlay UI (visible for 3 seconds after slide change)
		if manager.ShouldShowOverlay(3.0) {
			// Slide counter
			counter := fmt.Sprintf("%d / %d", manager.CurrentIndex()+1, manager.Count())
			rl.DrawText(counter, windowWidth-100, windowHeight-40, 20, rl.White)

			// Navigation hints
			hints := "Arrow Keys: Navigate | ESC: Exit"
			hintsWidth := rl.MeasureText(hints, 20)
			rl.DrawText(hints, (windowWidth-hintsWidth)/2, windowHeight-60, 20, rl.LightGray)
		}

		rl.EndDrawing()
	}
}

// InitialScreen displays a background image.
type InitialScreen struct {
	backgroundPath string
	background     rl.Texture2D
	pcTextures     [4]rl.Texture2D
	showPC         bool
	animatingPC    bool
	fadeStart      float64
	fadeDuration   float64
	frameAnimStart float64
	frameDuration  float64
}

func (s *InitialScreen) Load() {
	s.background = rl.LoadTexture(s.backgroundPath)
	s.pcTextures[0] = rl.LoadTexture("assets/pc/pc0.png")
	s.pcTextures[1] = rl.LoadTexture("assets/pc/pc1.png")
	s.pcTextures[2] = rl.LoadTexture("assets/pc/pc2.png")
	s.pcTextures[3] = rl.LoadTexture("assets/pc/pc3.png")
	s.showPC = false
	s.animatingPC = false
	s.fadeDuration = 1.0   // 1 second fade-in
	s.frameDuration = 0.25 // 0.25 seconds per frame
}

func (s *InitialScreen) Unload() {
	rl.UnloadTexture(s.background)
	for i := 0; i < 4; i++ {
		rl.UnloadTexture(s.pcTextures[i])
	}
}

func (s *InitialScreen) Update() screen.Screen {
	if rl.IsKeyPressed(rl.KeyEnter) {
		if !s.showPC {
			// First Enter: start fade-in
			s.showPC = true
			s.fadeStart = rl.GetTime()
		} else if !s.animatingPC {
			// Second Enter: start frame animation
			s.animatingPC = true
			s.frameAnimStart = rl.GetTime()
		}
	}
	return nil
}

func (s *InitialScreen) Draw() {
	// Scale texture to cover entire screen
	src := rl.Rectangle{X: 0, Y: 0, Width: float32(s.background.Width), Height: float32(s.background.Height)}
	dst := rl.Rectangle{X: 0, Y: 0, Width: float32(windowWidth), Height: float32(windowHeight)}
	rl.DrawTexturePro(s.background, src, dst, rl.Vector2{X: 0, Y: 0}, 0, rl.White)

	if s.showPC {
		// Calculate fade alpha
		fadeElapsed := rl.GetTime() - s.fadeStart
		fadeProgress := fadeElapsed / s.fadeDuration
		if fadeProgress > 1.0 {
			fadeProgress = 1.0
		}
		alpha := uint8(fadeProgress * 255)

		// Determine which frame to show
		frameIndex := 0
		if s.animatingPC {
			elapsed := rl.GetTime() - s.frameAnimStart
			cycleTime := 4 * s.frameDuration // 1 second for full cycle
			positionInCycle := elapsed - float64(int(elapsed/cycleTime))*cycleTime
			frameIndex = int(positionInCycle / s.frameDuration)
			if frameIndex > 3 {
				frameIndex = 3
			}
		}

		texture := s.pcTextures[frameIndex]
		x := (windowWidth - texture.Width) / 2
		y := (windowHeight - texture.Height) / 2
		rl.DrawTexture(texture, x, y, rl.Color{R: 255, G: 255, B: 255, A: alpha})
	}
}

// ProgrammingScreen displays a centered background image at original size.
type ProgrammingScreen struct {
	backgroundPath string
	background     rl.Texture2D
	screenTextures [4]rl.Texture2D
	typingTextures [4]rl.Texture2D
	animating      bool
	frameAnimStart float64
	frameDuration  float64
}

func (s *ProgrammingScreen) Load() {
	s.background = rl.LoadTexture(s.backgroundPath)
	s.screenTextures[0] = rl.LoadTexture("assets/screen/screen0.png")
	s.screenTextures[1] = rl.LoadTexture("assets/screen/screen1.png")
	s.screenTextures[2] = rl.LoadTexture("assets/screen/screen2.png")
	s.screenTextures[3] = rl.LoadTexture("assets/screen/screen3.png")
	s.typingTextures[0] = rl.LoadTexture("assets/typing/typing0.png")
	s.typingTextures[1] = rl.LoadTexture("assets/typing/typing1.png")
	s.typingTextures[2] = rl.LoadTexture("assets/typing/typing2.png")
	s.typingTextures[3] = rl.LoadTexture("assets/typing/typing3.png")
	s.animating = false
	s.frameDuration = 0.25
}

func (s *ProgrammingScreen) Unload() {
	rl.UnloadTexture(s.background)
	for i := 0; i < 4; i++ {
		rl.UnloadTexture(s.screenTextures[i])
		rl.UnloadTexture(s.typingTextures[i])
	}
}

func (s *ProgrammingScreen) Update() screen.Screen {
	if rl.IsKeyPressed(rl.KeyEnter) && !s.animating {
		s.animating = true
		s.frameAnimStart = rl.GetTime()
	}
	return nil
}

func (s *ProgrammingScreen) Draw() {
	x := (windowWidth - int32(s.background.Width)) / 2
	y := (windowHeight - int32(s.background.Height)) / 2
	rl.DrawTexture(s.background, x, y, rl.White)

	if s.animating {
		elapsed := rl.GetTime() - s.frameAnimStart
		cycleTime := 4 * s.frameDuration
		positionInCycle := elapsed - float64(int(elapsed/cycleTime))*cycleTime
		frameIndex := int(positionInCycle / s.frameDuration)
		if frameIndex > 3 {
			frameIndex = 3
		}

		texture := s.screenTextures[frameIndex]
		tx := int32(952)
		ty := int32(150)
		rl.DrawTexture(texture, tx, ty, rl.White)

		typingTexture := s.typingTextures[frameIndex]
		ttx := int32(790)
		tty := int32(700)
		rl.DrawTexture(typingTexture, ttx, tty, rl.White)
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
}
