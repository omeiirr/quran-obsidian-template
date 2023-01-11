package models

type Surah []struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Transliteration string `json:"transliteration"`
	Translation     string `json:"translation"`
	Type            string `json:"type"`
	TotalVerses     int    `json:"total_verses"`
	Verses          []Ayat `json:"verses"`
}
