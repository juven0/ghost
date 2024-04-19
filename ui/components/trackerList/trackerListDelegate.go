package trackerlist

import (
	"fmt"
	"ghost/ui/style"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ItemDelegate struct {
	//listMap *map[string]bool
}

func (d ItemDelegate) Height() int {
	return 4
}

func (d ItemDelegate) Spacing() int {
	return 0
}

func (d ItemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

func (d ItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	item, ok := listItem.(Item)
	if !ok {
		return
	}
	var trackTitle string
	trackTitle += style.TrackTitleStyle.Copy().Strikethrough(true).Render(item.title)
	if index == m.Index() {
		fmt.Fprint(w, style.TrackListActiveStyle.MaxWidth(m.Width()).Render(trackTitle))
	} else {
		fmt.Fprint(w, style.TrackListStyle.MaxWidth(m.Width()).Render(trackTitle))
	}
}
