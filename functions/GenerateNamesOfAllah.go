package functions

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/omeiirr/quran-obsidian-template/models"
	"github.com/omeiirr/quran-obsidian-template/templates"
)

func GenerateNamesOfAllah(workingDir string, embedded_file []byte) {
	fmt.Println("Generating names of Allah...")
	var names []models.NameOfAllah

	err := json.Unmarshal(embedded_file, &names)
	if err != nil {
		log.Fatal("Error during unmarshall: ", err)
	}

	file, err := os.OpenFile(fmt.Sprintf("%s/Names of Allah.md", workingDir), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
