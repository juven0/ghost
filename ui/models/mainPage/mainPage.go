package mainpage

import (
	trackerlist "ghost/ui/components/trackerList"
	"ghost/ui/style"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	width, height int
	program       *tea.Program
	trackerlist   *trackerlist.Model
	loaded        bool
	//tracker     *tracker.Model
}

// Init implements tea.Model.
func (m *Model) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.resize(msg.Width, msg.Height)
			m.trackerlist.InitList()
			m.loaded = true
			return m, tea.ClearScreen

		}
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	return m, cmd
}

// View implements tea.Model.
func (m *Model) View() string {
	if m.loaded {
		//return m.trackerlist.View()
		return style.TrackBoxStyle.Render(m.trackerlist.View())
	} else {
		return "loading..."
	}

	//return lipgloss.JoinVertical(lipgloss.Left, style.TrackBoxStyle.Render(m.trackerlist.View()), m.tracker.View())
}

func New() *Model {
	m := &Model{}
	p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseCellMotion())
	m.program = p
	m.loaded = false
	m.trackerlist = trackerlist.New(m.program)
	//m.tracker = tracker.New(p)
	return m
}

func (m *Model) InitLoad() {
	m.trackerlist.InitList()
}

func (m *Model) Run() {
	_, err := m.program.Run()
	if err != nil {
		return
	}
}

func (m *Model) resize(width, height int) {
	m.width, m.height = width, height
	//m.trackerlist.SetSize(m.width-4, height-14)
	//m.tracker.SetWidth(m.width - m.playlist.Width() - 4)

}
