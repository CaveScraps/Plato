package main

import (
	"os"
)

func CreateFileIfNotExists(filename string) error {
	if !fileExists(filename) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func prependToFile(filename string, data string) error {
	// read the existing file content
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// create a temporary file to write the modified content
	tempFile, err := os.CreateTemp(".", "")
	if err != nil {
		return err
	}

	// write the new line to the temp file then paste the original content.
	tempFile.WriteString(data + "\n")
	tempFile.WriteString(string(content))
	tempFile.Close()

	// replace the original file with the modified temp file
	err = os.Rename(tempFile.Name(), filename)
	if err != nil {
		return err
	}

	return nil
}

func appendToFile(filename string, data string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data + "\n")
	if err != nil {
		return err
	}

	return nil
}
