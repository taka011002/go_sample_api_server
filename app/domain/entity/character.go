package entity

type Character struct {
	Id    int `json:"id"`
	Name string `json:"name"`
	CharacterRarityId int `json:"character_rarity_id"`
	Power int `json:"power"`
}
