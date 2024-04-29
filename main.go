package main

import (
	trackerlist "ghost/ui/components/trackerList"

	tea "github.com/charmbracelet/bubbletea"
)

//"ghost/ui"

func main() {
	//ui.Run()
	m := &trackerlist.Model{}
	m.InitList()
	p := tea.NewProgram(m)

	_, err := p.Run()
	if err != nil {
		return
	}
}
