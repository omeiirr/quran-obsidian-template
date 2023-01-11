package models

type NameOfAllah struct {
	Arabic       string   `json:"arabic"`
	ArabicMSA    string   `json:"arabic_msa"`
	EnglishRoman string   `json:"english_roman"`
	English      []string `json:"english"`
	Reference    []string `json:"reference"`
	EnglishInfo  string   `json:"english_info"`
	Urdu         string   `json:"urdu"`
}
