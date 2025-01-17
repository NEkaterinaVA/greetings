package greetings

import (
    "errors"
    "fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string, randNum int) (string, error) {
    if name == "" {
        return "", errors.New("empty name")
    }
    message := fmt.Sprintf(randomFormat(randNum), name)
    return message, nil
}

func randomFormat(randNum int) string {
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }
    return formats[randNum]
}