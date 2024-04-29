package tracker

import (
	trackerlist "ghost/ui/components/trackerList"
	"ghost/ui/style"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	width    int
	program  *tea.Program
	progress progress.Model
	track    *trackerlist.Item
}

func New(p *tea.Program) *Model {
	m := &Model{
		progress: progress.New(progress.WithSolidFill(string(style.AccentColor))),
		program:  p,
		track:    &trackerlist.Item{},
	}
	m.progress.ShowPercentage = false
	m.progress.Empty = m.progress.Full
	m.progress.EmptyColor = string(style.BackgroundColor)
	return m
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) View() string {
	var playButton string
	// is playing ?
	playButton = style.ActiveButtonStyle.Padding(0, 1).Margin(0).Render(style.IconPlay)
	var trackTitle string
	trackTitle = style.TrackTitleStyle.Render(m.track.Title())
	trackTime := style.TrackVersionStyle.Render("4:32")
	trackAddInfo := style.TrackAddInfoStyle.Render(trackTime)
	trackTitle = lipgloss.JoinHorizontal(lipgloss.Top, trackTitle, trackAddInfo)

	tracker := style.TrackProgressStyle.Render(m.progress.View())
	tracker = lipgloss.JoinHorizontal(lipgloss.Top, playButton, tracker)
	tracker = lipgloss.JoinVertical(lipgloss.Left, tracker, trackTitle)

	return style.TrackBoxStyle.Width(m.width).Render(tracker)
}

func (m *Model) Update() (*Model, tea.Cmd) {
	return m, nil
}

func (m *Model) SetWidth(width int) {
	m.width = width
	m.progress.Width = width - 9
}

func (m *Model) Width() int {
	return m.width
}
