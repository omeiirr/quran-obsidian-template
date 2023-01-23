package functions

import (
	"fmt"
	"log"
	"os"
)

func GenerateFolderStructure(workingDir string) {
	fmt.Println("Generating folder structure...")
	// Create main folder

	err := os.MkdirAll(workingDir, 0750)
	if err != nil {
		log.Fatalf("Failed to create vault folder: %s", err)
	}

	// Additional sub-folders for better organisation
	directories := []string{""}

	for _, dir := range directories {
		filepath := fmt.Sprintf("%s/%s", workingDir, dir)
		err = os.MkdirAll(filepath, 0750)
		if err != nil {
			log.Fatalf("Failed to create sub folder: %s", err)
		}
	}

}
