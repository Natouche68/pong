package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	screenWidth  int
	screenHeight int
	grid         [][]string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.screenWidth = msg.Width
		m.screenHeight = msg.Height
		m.grid = make([][]string, m.screenHeight)
		for i := range m.grid {
			m.grid[i] = make([]string, m.screenWidth)
			for j := range m.grid[i] {
				m.grid[i][j] = " "
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	s := ""

	for i := range m.grid {
		for j := range m.grid[i] {
			s += m.grid[i][j]
		}
		s += "\n"
	}

	return s
}

func main() {
	p := tea.NewProgram(Model{}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program :", err)
		os.Exit(1)
	}
}
