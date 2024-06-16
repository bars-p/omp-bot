package logistic

import "fmt"

type Track struct {
	Title string
}

func (t *Track) String() string {
	return fmt.Sprintf("Package: %s", t.Title)
}
