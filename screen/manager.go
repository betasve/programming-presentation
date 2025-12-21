package screen

// Manager handles screen navigation and lifecycle.
type Manager struct {
	screens      []Screen
	currentIndex int
}

// NewManager creates a new screen manager.
func NewManager() *Manager {
	return &Manager{
		screens:      make([]Screen, 0),
		currentIndex: 0,
	}
}

// Add registers a screen to the presentation.
func (m *Manager) Add(s Screen) {
	m.screens = append(m.screens, s)
}

// Current returns the currently active screen.
func (m *Manager) Current() Screen {
	if len(m.screens) == 0 {
		return nil
	}
	return m.screens[m.currentIndex]
}

// Next advances to the next screen.
func (m *Manager) Next() {
	if len(m.screens) == 0 {
		return
	}
	if m.currentIndex < len(m.screens)-1 {
		m.screens[m.currentIndex].Unload()
		m.currentIndex++
		m.screens[m.currentIndex].Load()
	}
}

// Previous goes back to the previous screen.
func (m *Manager) Previous() {
	if len(m.screens) == 0 {
		return
	}
	if m.currentIndex > 0 {
		m.screens[m.currentIndex].Unload()
		m.currentIndex--
		m.screens[m.currentIndex].Load()
	}
}

// GoTo jumps to a specific screen by index.
func (m *Manager) GoTo(index int) {
	if len(m.screens) == 0 || index < 0 || index >= len(m.screens) {
		return
	}
	if index != m.currentIndex {
		m.screens[m.currentIndex].Unload()
		m.currentIndex = index
		m.screens[m.currentIndex].Load()
	}
}

// Count returns the total number of screens.
func (m *Manager) Count() int {
	return len(m.screens)
}

// CurrentIndex returns the index of the current screen.
func (m *Manager) CurrentIndex() int {
	return m.currentIndex
}

// Start initializes the presentation by loading the first screen.
func (m *Manager) Start() {
	if len(m.screens) > 0 {
		m.screens[0].Load()
	}
}

// Update calls Update on the current screen.
func (m *Manager) Update() {
	current := m.Current()
	if current != nil {
		if next := current.Update(); next != nil {
			// Screen requested a transition
			for i, s := range m.screens {
				if s == next {
					m.GoTo(i)
					return
				}
			}
		}
	}
}

// Draw calls Draw on the current screen.
func (m *Manager) Draw() {
	current := m.Current()
	if current != nil {
		current.Draw()
	}
}
