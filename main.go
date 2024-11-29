package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	screenWidth  int
	screenHeight int
	ball         Ball
}

type Ball struct {
	x         int
	y         int
	xVelocity int
	yVelocity int
}

type TickMsg struct{}

func doTick() tea.Cmd {
	return tea.Tick(time.Second/30, func(t time.Time) tea.Msg {
		return TickMsg{}
	})
}

func (m Model) Init() tea.Cmd {
	return doTick()
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

	case TickMsg:
		m.ball.x += m.ball.xVelocity
		m.ball.y += m.ball.yVelocity
		if m.ball.x <= 0 || m.ball.x >= m.screenWidth-1 {
			m.ball.xVelocity *= -1
		}
		if m.ball.y <= 0 || m.ball.y >= m.screenHeight-1 {
			m.ball.yVelocity *= -1
		}
		return m, doTick()
	}

	return m, nil
}

func (m Model) View() string {
	if m.screenWidth == 0 || m.screenHeight == 0 {
		return ""
	}

	grid := make([][]string, m.screenHeight)
	for i := range grid {
		grid[i] = make([]string, m.screenWidth)
		for j := range grid[i] {
			grid[i][j] = " "
		}
	}

	grid[m.ball.y][m.ball.x] = "@"

	s := ""
	for i := range grid {
		for j := range grid[i] {
			s += grid[i][j]
		}
		s += "\n"
	}

	return s
}

func main() {
	p := tea.NewProgram(Model{
		ball: Ball{
			x:         0,
			y:         0,
			xVelocity: 1,
			yVelocity: 1,
		},
	}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program :", err)
		os.Exit(1)
	}
}
