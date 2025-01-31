package domain

type Realm struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Flag       int    `json:"flag"`
	Icon       int    `json:"icon"`
	Population int    `json:"population"`
}

type Character struct {
	Guid  int    `json:"guid"`
	Name  string `json:"name"`
	Race  string `json:"race"`
	Class string `json:"class"`
	Level string `json:"level"`
}
