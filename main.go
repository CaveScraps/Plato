package main

import (
	"errors"
	"fmt"
    "os"
	"time"
)

type Config struct {
    date string;
    time string;
    message string;
}

// setup() returns a Config struct with the date, time, and message from the command line arguments.
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

func main() {
    config, err := setup()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    convertedDateInput, err := getDateFromInput(config.date)
    filename := convertedDateInput + ".md"

    if config.message == "" {
        fmt.Println("No message provided, exiting.")
        os.Exit(1)
    }

    var journalFile *os.File
    //If file doesn't exists already, then create it, otherwise open it to append.
    if _, err := os.Stat(filename); err != nil {
        file, err := os.Create(filename)
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
        journalFile = file
    }else{
        file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
        journalFile = file
    }

    //Append string to file
    defer journalFile.Close()
    _, err = journalFile.WriteString("\n" + config.time + ": " + config.message)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
