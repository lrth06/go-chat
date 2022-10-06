package structs
type Room struct {
	Id    string `json:"id"`
	Users []User `json:"users"`
}
