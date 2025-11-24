package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	Date  // embedding
}

func (e *Event) Title() string { // getter
	return e.title
}
func (e *Event) SetTitle(title string) error { // setter
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}
