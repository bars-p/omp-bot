package logistic

import "fmt"

type Track struct {
	ID uint64
	Title string
}

func (t *Track) String() string {
	return fmt.Sprintf("Track: [%d] %s", t.ID, t.Title)
}

func NewTrack(id uint64, title string) Track {
	return Track{
		ID: id,
		Title: title,
	}
}