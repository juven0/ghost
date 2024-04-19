package trackerlist

import (
	"fmt"
	"ghost/internal/songs"
	"ghost/ui/style"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const ROOT_DIR = "./music"

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type Model struct {
	list list.Model
}

func (m Model) Init() tea.Cmd {
	ch := make(chan []songs.Song)
	go songs.GetSongList(ROOT_DIR, ch)
	allSong := <-ch
	fmt.Println(allSong[0].Info.Name())
	return nil
}

func New() *Model {
	m := &Model{}
	m.initList()
	return m
}

func (m *Model) initList() {
	ch := make(chan []songs.Song)
	go songs.GetSongList(ROOT_DIR, ch)
	allSong := <-ch

	m.list = list.New([]list.Item{}, list.NewDefaultDelegate(), 1000, 1000)
	for i, song := range allSong {
		m.list.InsertItem(i, Item{title: song.Info.Name(), description: song.Path, Path: song.Path, IsPlaying: false})
	}
	// fmt.Printf(allSong[0].Info.Name())
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return style.FocusedStyle.Render(m.list.View())

}
