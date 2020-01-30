package entity

type ApiRequestLog struct {
	Id    int `json:"id"`
	UserId int `json:"user_id"`
	Method string `json:"method"`
	Path string `json:"path"`
	Params string `json:"params"`
}