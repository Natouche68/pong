package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
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
	}

	return m, nil
}

func (m Model) View() string {
	return "Pong"
}

func main() {
	p := tea.NewProgram(Model{})
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program :", err)
		os.Exit(1)
	}
}
