package mainpage

import trackerlist "ghost/ui/components/trackerList"

type Model struct {
	trackerlist *trackerlist.Model
}

func New() *Model {
	m := &Model{}
	m.trackerlist = trackerlist.New()
	return m
}

func (m *Model) Run() {
	m.trackerlist.Init()
}

// func (m *Model) Init() void {

// }

// func (m *Model) Update() void {

// }

// func (m *Model) View() string {

// }
