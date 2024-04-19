package trackerlist

import (
<<<<<<< HEAD
	"fmt"
	"ghost/internal/songs"

=======
>>>>>>> c4dae7c705aa1e716e82b618fd136619b03f4e76
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

<<<<<<< HEAD
const ROOT_DIR = "./music"

=======
>>>>>>> c4dae7c705aa1e716e82b618fd136619b03f4e76
var docStyle = lipgloss.NewStyle().Margin(1, 2)

type Model struct {
	list list.Model
}

func (m Model) Init() tea.Cmd {
<<<<<<< HEAD
	ch := make(chan []songs.Song)
	go songs.GetSongList(ROOT_DIR, ch)
	allSong := <-ch
	fmt.Println(allSong[0].Info.Name())
=======
>>>>>>> c4dae7c705aa1e716e82b618fd136619b03f4e76
	return nil
}

func New() *Model {
	m := &Model{}
<<<<<<< HEAD
	m.initList()
	return m
}

func (m *Model) initList() {
	ch := make(chan []songs.Song)
	go songs.GetSongList(ROOT_DIR, ch)
	allSong := <-ch

	m.list = list.New([]list.Item{}, list.NewDefaultDelegate(), 500, 500)
	for i, song := range allSong {
		m.list.InsertItem(i, Item{title: song.Info.Name(), description: song.Info.Name(), Path: song.Path, IsPlaying: false})
	}
	// fmt.Printf(allSong[0].Info.Name())
}

=======
	//***
	//get and the list of songs
	//***
	return m
}

>>>>>>> c4dae7c705aa1e716e82b618fd136619b03f4e76
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
	return docStyle.Render(m.list.View())
}
