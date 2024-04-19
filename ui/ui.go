package ui

import mainpage "ghost/ui/models/mainPage"

func Run() {
	var err error

	mainpage.New().Run()
	if err != nil {

	}
}
