package app

type Gameroom struct {
	Id       int    `json:"-"`
	Creator  string `json:"creator"`
	Capacity int    `json:"-"`
	Visitors []string
}
