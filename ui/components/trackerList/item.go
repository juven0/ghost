package trackerlist

type Item struct {
	Name      string
	Artist    string
	Path      string
	IsPlaying bool
}

func (i Item) FilterValue() string {
	return i.Name
}
