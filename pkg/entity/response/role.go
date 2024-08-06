package response

type RoleResponse struct {
    Code string
    Message string
    Roles []Role
}

type Role struct {
    ID          uint 
    UUID        string 
    CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}