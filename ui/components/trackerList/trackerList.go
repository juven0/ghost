package trackerlist

import (
	"errors"
	"fmt"
	"ghost/internal/audio"
	"ghost/internal/songs"
	"ghost/ui/style"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const ROOT_DIR = "./music"

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type Model struct {
	width, height int
	program       *tea.Program
	list          list.Model
	cursor        int
	audioPanel    *audio.AudioPanel
}

func (m Model) Init() tea.Cmd {
	ch := make(chan []songs.Song)
	go songs.GetSongList(ROOT_DIR, ch)
	allSong := <-ch
	fmt.Println(allSong[0].Info.Name())
	m.InitList()
	return nil
}

func New(p *tea.Program) *Model {
	m := &Model{
		program: p,
		cursor:  0,
	}
	return m
}

func (m *Model) InitList() {
	ch := make(chan []songs.Song)
	go songs.GetSongList(ROOT_DIR, ch)
	allSong := <-ch

	m.list = list.New([]list.Item{}, ItemDelegate{}, 1000, 1000)
	for i, song := range allSong {
		m.list.InsertItem(i, Item{title: song.Info.Name(), description: song.Path, Path: song.Path, IsPlaying: false})
	}
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			item, err := m.getItems(m.cursor)
			if err != nil {

			}
			ap := &audio.AudioPanel{}
			ap.Play(item.Path)
			m.audioPanel = ap
		}
		if msg.String() == "space" {
			if m.audioPanel != nil {
				m.audioPanel.Stop()
				m.audioPanel = nil
			}
		}
		if msg.String() == "up" {
			if m.cursor > 0 {
				m.cursor -= 1
			}
		}
		if msg.String() == "down" {
			if m.cursor < len(m.list.Items()) {
				m.cursor += 1
			}
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	fmt.Println(m.cursor)
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *Model) View() string {
	return style.FocusedStyle.Render(m.list.View())

}

func (m *Model) getItems(pos int) (Item, error) {
	for index, item := range m.list.Items() {
		entry, ok := item.(Item)
		if !ok {
			continue
		}
		if index == pos {
			return entry, nil
		}
	}
	return Item{}, errors.New("null")
}

func (m *Model) SetSize(w, h int) {
	m.width = w
	m.height = h
	m.list.SetSize(m.width, m.height)
}

func (m *Model) SetWidth(w int) {
	m.width = w
	m.list.SetSize(m.width, m.height)
}

func (m *Model) Width() int {
	return m.width
}

func (m *Model) SetHeight(h int) {
	m.height = h
	m.list.SetSize(m.width, m.height)
}

func (m *Model) Height() int {
	return m.height
}
