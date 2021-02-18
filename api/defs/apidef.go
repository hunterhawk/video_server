package defs

// requests
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}

// Data model

type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}