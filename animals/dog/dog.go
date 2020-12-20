package dog

import (  
    "fmt"
    "errors"
)

type Dog struct {  
    Name string
}

func (a Dog) Speak() (string, error) { 
	if a.Name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("I am a dog, my name is %s", a.Name)
	return message, nil
}