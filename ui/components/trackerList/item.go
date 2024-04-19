package trackerlist

type Item struct {
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
}
