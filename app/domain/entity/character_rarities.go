package entity

type CharacterRarity struct {
	Id    int `json:"id"`
	Name string `json:"name"`
	Rarity int `json:"rarity"`
}

type CharacterRarities []CharacterRarity

func (il CharacterRarities) Len() int           { return len(il) }
func (il CharacterRarities) Less(i, j int) bool { return il[i].Rarity < il[j].Rarity }
func (il CharacterRarities) Swap(i, j int)      { il[i], il[j] = il[j], il[i] }