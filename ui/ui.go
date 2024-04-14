package ui

import mainpage "ghost/ui/models/mainPage"

func Run() {
	var err error

	err = mainpage.New().Run()
	if err != nil {

	}
}
