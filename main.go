package main

import (
	trackerlist "ghost/ui/components/trackerList"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// ui.Run()
	m := trackerlist.New()
	p := tea.NewProgram(m)

	if err := p.Start(); err != nil {
		os.Exit(1)
	}

}
