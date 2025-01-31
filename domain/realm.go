package domain

type Realm struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Flag       int    `json:"flag"`
	Icon       int    `json:"icon"`
	Population int    `json:"population"`
}
