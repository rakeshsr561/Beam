package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFileByLine(fileName string) []string {
	urls := []string{}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		urls = append(urls, string(scanner.Text()))

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return urls
}

func CreateFolderIfNotExists(path string) error {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, 0700)
		return err
	}
	return nil
}
