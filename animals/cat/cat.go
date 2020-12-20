package cat

import (  
    "fmt"
    "errors"
)

type Cat struct {  
    Name string
}

func (a Cat) Speak() (string, error) { 
	if a.Name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("I am a cat, my name is %s", a.Name)
	return message, nil
}