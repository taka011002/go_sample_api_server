package entity

type Ranking struct {
	Rank  int  `json:"rank"`
	User  User `json:"user"`
	Score int  `json:"score"`
}
