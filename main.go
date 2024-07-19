package main

import (
	"errors"
	"fmt"
    "os"
	"time"
    "flag"
)

type Config struct {
    date string
    time string
    message string
    isTodoItem bool
}

func setup() (Config, error){
    todoPtr := flag.Bool("t", false, "Add a todo item instead of a journal entry")
    flag.Parse()
    os.Args = flag.Args()

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

    return Config{date: date, time: time, message: message, isTodoItem: *todoPtr}, nil
}


func main() {
    config, err := setup()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    convertedDateInput, err := getDateFromInput(config.date)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    if config.message == "" {
        fmt.Println("No message provided, exiting.")
        os.Exit(1)
    }

    filename := convertedDateInput + ".md"
    if err := CreateFileIfNotExists(filename); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    //Todos go to the bottom of the file, the rest is placed newest at the top.
    if config.isTodoItem {
        message := "- [ ] " + config.message
        err := appendToFile(filename, message)
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
    }else{
        message := config.time + " " + config.message
        err := prependToFile(filename, message)
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
    }
}
