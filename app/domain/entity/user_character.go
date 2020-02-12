package entity

type UserCharacter struct {
	Id          int       `json:"id"`
	UserId      string    `json:"user_id"`
	CharacterId string    `json:"character_id"`
	Character   Character `json:"character"`
}

type UserCharacters []UserCharacter
