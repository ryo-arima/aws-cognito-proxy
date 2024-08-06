package response

type PolicyResponse struct {
    Code string
    Message string
    Policys []Policy
}

type Policy struct {
    ID          uint 
    UUID        string 
    CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}