package main

import (
<<<<<<< HEAD
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

=======
	"ghost/ui"
)

func main() {
	ui.Run()
>>>>>>> c4dae7c705aa1e716e82b618fd136619b03f4e76
}
