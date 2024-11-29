package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	gameStarted  bool
	gameOver     bool
	screenWidth  int
	screenHeight int
	ball         Ball
	leftPad      Pad
	rightPad     Pad
}

type Ball struct {
	x         int
	y         int
	xVelocity int
	yVelocity int
}

type Pad struct {
	y         int
	yVelocity int
	size      int
}

type TickMsg struct{}

func doTick() tea.Cmd {
	return tea.Tick(time.Second/16, func(t time.Time) tea.Msg {
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

		case "enter":
			if !m.gameStarted {
				m.gameStarted = true
				m.ball = Ball{
					x:         0,
					y:         0,
					xVelocity: 2,
					yVelocity: 1,
				}
				m.leftPad = Pad{
					y:         0,
					yVelocity: 1,
					size:      m.screenHeight / 4,
				}
				m.rightPad = Pad{
					y:         0,
					yVelocity: 1,
					size:      m.screenHeight / 4,
				}
			}

		case "z", "w":
			m.leftPad.yVelocity = -1

		case "s":
			m.leftPad.yVelocity = 1

		case "up":
			m.rightPad.yVelocity = -1

		case "down":
			m.rightPad.yVelocity = 1
		}

	case tea.WindowSizeMsg:
		m.screenWidth = msg.Width
		m.screenHeight = msg.Height

	case TickMsg:
		if !m.gameStarted {
			return m, doTick()
		}

		m.ball.x += m.ball.xVelocity
		m.ball.y += m.ball.yVelocity
		if m.ball.x <= 0 {
			m.ball.xVelocity *= -1
			if m.ball.y < m.leftPad.y || m.ball.y >= m.leftPad.y+m.leftPad.size {
				m.gameOver = true
				m.gameStarted = false
			}
		}
		if m.ball.x >= m.screenWidth-2 {
			m.ball.xVelocity *= -1
			if m.ball.y < m.rightPad.y || m.ball.y >= m.rightPad.y+m.rightPad.size {
				m.gameOver = true
				m.gameStarted = false
			}
		}
		if m.ball.y <= 0 || m.ball.y >= m.screenHeight-2 {
			m.ball.yVelocity *= -1
		}

		if (m.leftPad.yVelocity == -1 && m.leftPad.y > 0) || (m.leftPad.yVelocity == 1 && m.leftPad.y < m.screenHeight-m.leftPad.size) {
			m.leftPad.y += m.leftPad.yVelocity
		}
		if (m.rightPad.yVelocity == -1 && m.rightPad.y > 0) || (m.rightPad.yVelocity == 1 && m.rightPad.y < m.screenHeight-m.rightPad.size) {
			m.rightPad.y += m.rightPad.yVelocity
		}

		return m, doTick()
	}

	return m, nil
}

func (m Model) View() string {
	if !m.gameStarted {
		if m.gameOver {
			return "GAME OVER !\nPress ENTER to restart..."
		} else {
			return "Press ENTER to start game..."
		}
	}

	grid := make([][]string, m.screenHeight)
	for i := range grid {
		grid[i] = make([]string, m.screenWidth)
		for j := range grid[i] {
			grid[i][j] = " "
		}
	}

	grid[m.ball.y][m.ball.x] = "@"

	for i := 0; i < m.leftPad.size; i++ {
		grid[m.leftPad.y+i][0] = "#"
	}
	for i := 0; i < m.rightPad.size; i++ {
		grid[m.rightPad.y+i][m.screenWidth-1] = "#"
	}

	s := ""
	for i := range grid {
		for j := range grid[i] {
			s += grid[i][j]
		}
		if i < m.screenHeight-1 {
			s += "\n"
		}
	}

	return s
}

func main() {
	p := tea.NewProgram(Model{
		gameStarted: false,
	}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program :", err)
		os.Exit(1)
	}
}
