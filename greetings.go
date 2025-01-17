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

func Hellos(names []string, randNum int)(map[string]string, error){
    messages := make(map[string]string)
    for _, name := range names {
        message, err := Hello(name, randNum)
        if err != nil {
            return nil, err
        }
        messages[name] = message
    }
    return messages, nil
}

func randomFormat(randNum int) string {
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }
    return formats[randNum]
}