package entity

type User struct {
	Id             int              `json:"id"`
	Username       string           `json:"username"`
	Password       string           `json:"password"`
	UserCharacters []UserCharacters `json:"user_characters"`
}
