package models

type Ayat struct {
	Id          int    `json:"id"`
	Text        string `json:"text"`
	Translation string `json:"translation"`
}
