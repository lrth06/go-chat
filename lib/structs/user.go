package structs
type User struct {
	ID     any           `json:"_id"`
	Name   string        `json:"username"`
	Email  string        `json:"email" validate:"required,email"`
	Avatar string        `json:"avatar" validate:"omitempty,url"`
	Token  string        `json:"token"`
	Roles  []interface{} `json:"roles"`
}
