package barkbehaviour

import (  
    "fmt"
)

type SpeakBehaviour interface {
	Speak()
}

func Speak() string {  
    return fmt.Sprintf("Woff!")
}