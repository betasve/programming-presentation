package input

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/betasve/programming-presentation/screen"
)

// Handle processes input and updates the screen manager accordingly.
func Handle(m *screen.Manager) {
	// Navigation: Next slide
	if rl.IsKeyPressed(rl.KeyRight) || rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyEnter) {
		m.Next()
	}

	// Navigation: Previous slide
	if rl.IsKeyPressed(rl.KeyLeft) || rl.IsKeyPressed(rl.KeyBackspace) {
		m.Previous()
	}

	// Toggle fullscreen
	if rl.IsKeyPressed(rl.KeyF) {
		rl.ToggleFullscreen()
	}

	// Number keys for direct navigation (1-9 for slides 0-8)
	for i := int32(rl.KeyOne); i <= int32(rl.KeyNine); i++ {
		if rl.IsKeyPressed(i) {
			slideIndex := int(i - int32(rl.KeyOne))
			m.GoTo(slideIndex)
		}
	}

	// Zero key for slide 10 (index 9)
	if rl.IsKeyPressed(rl.KeyZero) {
		m.GoTo(9)
	}
}
