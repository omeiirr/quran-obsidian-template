package functions

import (
	"fmt"
	"log"
	"os"
)

func GenerateFolderStructure() {
	// Create main folder
	err := os.MkdirAll("Quranic Guidance", 0750)
	if err != nil {
		log.Fatalf("Failed to create vault folder: %s", err)
	}

	// Additional sub-folders for better organisation
	directories := [3]string{"People/Prophets", "Themes"}

	for _, dir := range directories {
		filepath := fmt.Sprintf("./Quranic Guidance/%s", dir)
		err = os.MkdirAll(filepath, 0750)
		if err != nil {
			log.Fatalf("Failed to create sub folder: %s", err)
		}
	}

}
