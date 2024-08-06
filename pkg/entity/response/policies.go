package response

import "time"

type PoliciesResponse struct {
	Code      string
	Message   string
	Policiess []Policies
}

type Policies struct {
	ID        uint
	UUID      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
