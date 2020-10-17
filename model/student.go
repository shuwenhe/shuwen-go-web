package model

type Student struct {
	ID    int    `json:"id,omitempty"`
	Num   string `json:"num,omitempty"`
	Name  string `json:"name,omitempty"`
	Pass  string `json:"pass,omitempty"`
	Phone string `json:"phone,omitempty"`
}
