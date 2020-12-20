package human

import (  
    "fmt"
    "errors"
)

type Human struct {  
    Name string
}

func (a Human) Speak() (string, error) { 
	if a.Name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("I am a human, my name is %s", a.Name)
	return message, nil
}