package targets

import (
	"log"
	"os"
)

var (
	readmeFilename = "README.md"
)

const readmeContent = `# README FILE
`

func createReadme() error {
	if _, err := os.Stat(readmeFilename); os.IsNotExist(err) {
		f, err := os.Create(readmeFilename)
		if err != nil {
			return err
		}

		defer f.Close()

		_, err = f.WriteString(readmeContent)
		if err != nil {
			return err
		}
	} else {
		log.Printf("%s already exists. ignoring", readmeFilename)
	}

	return nil
}