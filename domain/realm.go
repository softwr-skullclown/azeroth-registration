package domain

type Realm struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Flag       int    `json:"flag"`
	Icon       int    `json:"icon"`
	Population int    `json:"population"`
}

type Character struct {
	Guid   int    `json:"guid"`
	Name   string `json:"name"`
	Race   int    `json:"race"`
	Class  int    `json:"class"`
	Level  int    `json:"level"`
	Gender int    `json:"gender"`
}
