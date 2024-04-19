package trackerlist

type Item struct {
<<<<<<< HEAD
	title       string
	description string
	Path        string
	IsPlaying   bool
}

func (i Item) FilterValue() string {
	return i.title
}

func (i Item) Title() string {
	return i.title
}

func (i Item) Description() string {
	return i.description
=======
	Name      string
	Artist    string
	Path      string
	IsPlaying bool
}

func (i Item) FilterValue() string {
	return i.Name
>>>>>>> c4dae7c705aa1e716e82b618fd136619b03f4e76
}
