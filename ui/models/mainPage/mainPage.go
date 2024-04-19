package mainpage

<<<<<<< HEAD
import trackerlist "ghost/ui/components/trackerList"
=======
import (
	"fmt"
	trackerlist "ghost/ui/components/trackerList"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)
>>>>>>> c4dae7c705aa1e716e82b618fd136619b03f4e76

type Model struct {
	trackerlist *trackerlist.Model
}

func New() *Model {
	m := &Model{}
	m.trackerlist = trackerlist.New()
	return m
}

<<<<<<< HEAD
func (m *Model) Run() {
	m.trackerlist.Init()
}

// func (m *Model) Init() void {

// }

// func (m *Model) Update() void {

// }

// func (m *Model) View() string {

// }
=======
func (m *Model) Run() error {
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
	return nil
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	return m, cmd
}

func (m *Model) View() string {
	return m.trackerlist.View()
}
>>>>>>> c4dae7c705aa1e716e82b618fd136619b03f4e76
