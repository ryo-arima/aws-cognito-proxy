package response

import "time"

type GroupsResponse struct {
	Code    string
	Message string
	Groupss []Groups
}

type Groups struct {
	ID        uint
	UUID      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
