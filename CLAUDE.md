# Programming Presentation

Interactive presentation software built with Go and raylib-go.

## Commands

```bash
# Run the presentation
go run .

# Build binary
go build -o presentation .

# Run with race detector (debugging)
go run -race .

# Format code
gofmt -w .

# Vet code
go vet ./...

# Tidy dependencies
go mod tidy
```

## Architecture

- `main.go` - Entry point, window init, game loop
- `screen/screen.go` - Screen interface that all slides implement
- `screen/manager.go` - Handles navigation between screens
- `input/input.go` - Keyboard input handling

## Creating New Screens

Implement the `screen.Screen` interface:

```go
type MyScreen struct{}

func (s *MyScreen) Load()   {}           // Called when screen becomes active
func (s *MyScreen) Unload() {}           // Called when leaving screen
func (s *MyScreen) Update() screen.Screen { return nil }  // Return nil to stay
func (s *MyScreen) Draw()   {}           // Render content
```

Register in main.go: `manager.Add(&MyScreen{})`

## Controls

- Right/Space/Enter: Next slide
- Left/Backspace: Previous slide
- 1-9, 0: Jump to slide
- F: Toggle fullscreen
- ESC: Exit
