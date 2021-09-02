package types

type Member struct {
	UUID               string   `json:"uuid"`
	CreatedTime        string   `json:"created_time"`
	Type               string   `json:"type" validate:"eq=contractor|eq=employee"`
	DurationOfContract int32    `json:"duration_of_contract"`
	Skills             []string `json:"skills"`
	MemberName         string   `json:"member_name" validate:"required"`
	Title              string   `json:"title" validate:"required"`
	Email              string   `json:"email" validate:"required,email"`
}
