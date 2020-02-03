package entity

type User struct {
	Id    int `json:"id"`
	Username string `json:"username"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Phone string `json:"phone"`
	UserStatus int `json:"user_status"`
	UserCharacters []UserCharacters `json:"user_characters"`
}