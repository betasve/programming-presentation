package screen

// Screen defines the interface that all presentation screens must implement.
type Screen interface {
	// Load is called once when the screen becomes active.
	Load()

	// Unload is called when leaving the screen.
	Unload()

	// Update handles screen logic and returns the next screen to transition to.
	// Return nil to stay on the current screen.
	Update() Screen

	// Draw renders the screen content.
	Draw()
}
