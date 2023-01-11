package functions

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/omeiirr/quran-obsidian-template/models"
	"github.com/omeiirr/quran-obsidian-template/templates"
)

func GenerateNamesOfAllah() {
	content, err := ioutil.ReadFile("./data/asma-ul-husna.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var names []models.NameOfAllah

	err = json.Unmarshal(content, &names)
	if err != nil {
		log.Fatal("Error during unmarshall: ", err)
	}

	file, err := os.OpenFile("./Quranic Guidance/Names of Allah.md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}

	writer := bufio.NewWriter(file)

	writer.WriteString(templates.NamesOfAllahTemplate)

	for _, name := range names {
		refs := strings.Join(name.Reference, ", ")
		writer.WriteString(fmt.Sprintf("\n| %s | %s | %s | %s | %v |", name.EnglishRoman, name.Arabic, name.English[0], name.EnglishInfo, strings.ReplaceAll(refs, ", ", "<br/>")))
	}
	writer.Flush()
	file.Close()

}
