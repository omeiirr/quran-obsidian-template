package functions

import (
	"encoding/json"
	"fmt"
	"log"

	"os"
	"path/filepath"
	"strings"

	"github.com/omeiirr/quran-obsidian-template/templates"
)

func GenerateOntology(workingDir string, embedded_file []byte) {

	var concepts []string

	err := json.Unmarshal(embedded_file, &concepts)
	if err != nil {
		log.Fatal("Error during unmarshall: ", err)
	}

	for _, concept := range concepts {

		// extract the directory path
		dir := filepath.Dir(concept)

		// create the directory, including any necessary parent directories
		err := os.MkdirAll(fmt.Sprintf("%s/%s", workingDir, dir), os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}

		conceptName := strings.Split(filepath.Base(concept), ".md")[0]
		urlKeyword := strings.Split(conceptName, " (")[0]

		// convert conceptName to lowercase, then replace ALL occurences of " " with "-"
		// in accordance with the URLs of `corpus.quran.com`
		externalUrl := fmt.Sprintf("https://corpus.quran.com/concept.jsp?id=%s", strings.Replace(strings.ToLower(urlKeyword), " ", "-", -1))

		// define the content for each concept and insert into the template
		content := []byte(
			fmt.Sprintf(templates.OntologyConceptTemplate, externalUrl))

		err = os.WriteFile(fmt.Sprintf("%s/%s", workingDir, concept), content, 0660)
		if err != nil {
			fmt.Println(err)
		}
	}
}
