package main

import (
	"image"
	"log"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ryan4yin/video2ascii/internal/video2ascii"
)

var cursor int

type model struct {
	charImgList []video2ascii.CharImg
	fps         float64
}

type tickMsg time.Time

func main() {
	const videoPath = "./test_data/BadApple.mp4"
	var size = image.Point{64, 48}
	const seconds = 30
	charImgList, fps, err := video2ascii.Video2Chars(videoPath, size, seconds)
	if err != nil {
		log.Fatalf("failed to convert Video to Chars: %v", err)
	}

	p := tea.NewProgram(model{charImgList, fps}, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(tick(m.fps), tea.EnterAltScreen)
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	case tickMsg:
		if cursor >= len(m.charImgList) {
			return m, tea.Quit
		}
		return m, tick(m.fps)

	}

	return m, nil
}

func (m model) View() string {
	s := strings.Join(m.charImgList[cursor], "\n")
	cursor++
	return s
}

func tick(fps float64) tea.Cmd {
	return tea.Tick(time.Duration(float64(time.Second)/fps), func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
