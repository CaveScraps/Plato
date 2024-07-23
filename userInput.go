package main

import (
	"errors"
	"os"
	"os/exec"
	"time"
)

// getDateFromInput takes a string and returns a string in the format "YYYY-MM-DD" or an error.
func getDateFromInput(dateString string) (string, error) {
	if dateString == "today" {
		return time.Now().Format("2006-01-02"), nil
	} else if dateString == "tomorrow" {
		return time.Now().AddDate(0, 0, 1).Format("2006-01-02"), nil
	} else if dateString == "yesterday" {
		return time.Now().AddDate(0, 0, -1).Format("2006-01-02"), nil
	}

	// Check if the date is in the correct format
	_, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return "", errors.New("Date is not in the correct format")
	}

	// If the date is not a special case and it is in the correct format, use it as is.
	return dateString, nil
}

// getInputFromVim opens vim for the user to input a message and returns the message as a string or an error.
func getInputFromVim(workingDir string) (string, error) {
	tempFile, err := os.CreateTemp(workingDir, "")
	if err != nil {
		return "", err
	}
	defer os.Remove(tempFile.Name())

	vimCommand := exec.Command("nvim", tempFile.Name())
	vimCommand.Stdin = os.Stdin
	vimCommand.Stdout = os.Stdout
	err = vimCommand.Run()
	if err != nil {
		return "", errors.New("Error running nvim")
	}

	data, err := os.ReadFile(tempFile.Name())
	if err != nil {
		return "", errors.New("Error reading nvim's output")
	}

	return string(data), nil
}
