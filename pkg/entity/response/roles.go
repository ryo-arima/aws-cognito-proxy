package response

import "time"

type RolesResponse struct {
	Code    string
	Message string
	Roless  []Roles
}

type Roles struct {
	ID        uint
	UUID      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
