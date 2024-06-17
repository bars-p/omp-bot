package logistic

import "fmt"

type Track struct {
	ID uint64
	Title string
}

func (t *Track) String() string {
	return fmt.Sprintf("Package: [%d] %s", t.ID, t.Title)
}
