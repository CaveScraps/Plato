package main

import (
	"os"
	"strings"
	"time"
)

func createJournalFileIfNotExists(date string, location string) error {
	filename := location + date + ".md"
	if !fileExists(filename) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}

		err = copyIncompleteTodos(location, date)
		if err != nil {
			return err
		}
	}
	return nil
}

func copyIncompleteTodos(noteDir string, currentDay string) error {
	currentDayAsDate, err := time.Parse("2006-01-02", currentDay)
	if err != nil {
		return err
	}

	previousDay := currentDayAsDate.AddDate(0, 0, -1)
	previousDayFile := noteDir + previousDay.Format("2006-01-02") + ".md"

	// If the previous day file does not exist, there is nothing to copy.
	if fileExists(previousDayFile) == false {
		return nil
	}

	content, err := os.ReadFile(previousDayFile)
	if err != nil {
		return err
	}

	currentDayFile := noteDir + currentDay + ".md"
	for _, line := range strings.Split(string(content), "\n") {
		if strings.HasPrefix(line, "- [ ]") {
			err := appendToFile(currentDayFile, line)
			if err != nil {
				return err
			}
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
