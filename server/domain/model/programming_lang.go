package model

import "time"

// ProgrammingLang は、プログラミング言語を表す。
type ProgrammingLang struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Feature   string    `json:"feature"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
