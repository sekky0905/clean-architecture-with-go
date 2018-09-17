package model

import "time"

// ProgrammingLang は、プログラミング言語を表す。
type ProgrammingLang struct {
	ID        int
	Name      string
	Feature   string
	CreatedAt time.Time
	UpdatedAt time.Time
}