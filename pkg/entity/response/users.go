package response

import "time"

type UsersResponse struct {
	Code    string
	Message string
	Userss  []Users
}

type Users struct {
	ID        uint
	UUID      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
