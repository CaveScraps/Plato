package main

import (
	"errors"
	"fmt"
    "os"
	"os/exec"
	"time"
)

type Config struct {
    date string;
    time string;
    message string;
}

func setup() (Config, error){
    os.Args = os.Args[1:]

    // Looks like golang flags doesn't allow flags after other args so for now we will use this.
    if len(os.Args) < 1 {
        return Config{}, errors.New("No arguments provided")
    }
    if len(os.Args) > 2 {
        return Config{}, errors.New("Too many arguments provided, is your message in quotes?")
    }

    date := string(os.Args[0])
    time := time.Now().Format("15:04")
    message := ""

    if len(os.Args) == 1 {
        // If there is no command line message, open vim for user to input message.
        data, err := getInputFromVim("temp.txt")
        message = data
        if err != nil {
            return Config{}, err
        }
    }else{
        message = string(os.Args[1])
    }

    return Config{date: date, time: time, message: message}, nil
}

func getDateFromInput(dateString string) (string, error) {
    if dateString == "today" {
        return time.Now().Format("2006-01-02"), nil
    }else if dateString == "tomorrow" {
        return time.Now().AddDate(0, 0, 1).Format("2006-01-02"), nil
    }else if dateString == "yesterday" {
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

func getInputFromVim(fileName string) (string, error) {
    os.Create("temp.txt")
    vimCommand := exec.Command("nvim", fileName)
    vimCommand.Stdin = os.Stdin
    vimCommand.Stdout = os.Stdout
    err := vimCommand.Run()
    if err != nil {
        return "", errors.New("Error running nvim")
    }

    data, err := os.ReadFile("temp.txt")
    if err != nil {
        return "", errors.New("Error reading nvims output")
    }

    return string(data), nil
}

func main() {
    config, err := setup()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println("Hello, World!\n" + config.date + "\n" + config.time + "\n" + config.message)
}
