package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/omeiirr/quran-obsidian-template/models"
	"github.com/omeiirr/quran-obsidian-template/templates"
)

func GenerateQuran() {

	res, err := http.Get("https://cdn.jsdelivr.net/npm/quran-json@3.1.2/dist/quran_en.json")

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Failed to read body of API response: %s", err)
	}

	surahs := models.Surah{}
	err = json.Unmarshal(body, &surahs)
	if err != nil {
		log.Printf("Failed to read body of Surah: %s", err)
		return
	}

	for _, surah := range surahs[113:] {

		// Create a folder for each Surah
		surahFilepath := fmt.Sprintf("./Quranic Guidance/%d %s (%s)", surah.Id, surah.Transliteration, surah.Translation)
		err = os.MkdirAll(surahFilepath, 0750)
		if err != nil {
			log.Fatalf("Failed to create Surah folders: %s", err)
		}

		fmt.Printf("\nTotal verses: %d \n", surah.TotalVerses)
		for _, ayat := range surah.Verses {
			fmt.Printf("%d ", ayat.Id)

			filename := fmt.Sprintf("%s/%d-%d.md", surahFilepath, surah.Id, ayat.Id)
			// content := []byte(verse.Text)
			content := []byte(
				fmt.Sprintf(templates.AyatTemplate, surah.Type, ayat.Text, ayat.Translation))

			err = os.WriteFile(filename, content, 0660)
			if err != nil {
				log.Fatalf("Failed to write Ayat into file: %s", err)
			}
		}
	}

}
