package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

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

	for _, surah := range surahs {

		// Create a folder for each Surah

		// Windows does not allow special characters in filename
		// Problem detected for surah-38 with Translation as `"The Letter \"Saad\""` from API
		filename := strings.ReplaceAll(surah.Translation, "\"", "'")

		surahFilepath := fmt.Sprintf("./Quranic Guidance/Surahs/%d %s (%s)", surah.Id, surah.Transliteration, filename)
		err = os.MkdirAll(surahFilepath, 0750)
		if err != nil {
			log.Fatalf("Failed to create Surah folders: %s", err)
		}

		fmt.Printf("\nGenerating Surah: %s", surah.Translation)
		for _, ayat := range surah.Verses {
			filename := fmt.Sprintf("%s/%d-%d.md", surahFilepath, surah.Id, ayat.Id)
			content := []byte(
				fmt.Sprintf(templates.AyatTemplate, surah.Type, ayat.Text, ayat.Translation))

			err = os.WriteFile(filename, content, 0660)
			if err != nil {
				log.Fatalf("Failed to write Ayat into file: %s", err)
			}
		}
	}

}

// res, err := http.Get("https://quranenc.com/api/v1/translation/sura/english_saheeh/1?translation_key=english_rwwad&sura_number=110")
