package main

import (  
    "fmt"
    "test.com/supertype"
    "test.com/dog"
    "test.com/cat"
    "test.com/human"
)

type Weird struct {  
    Name string
}


func main() {
    var animal supertype.SpeakBehaviour = dog.Dog{"Pepe dog"}
    fmt.Println(animal.Speak())

    animal = cat.Cat{"Meow cat"}
    fmt.Println(animal.Speak())

    animal = human.Human{"The Human"}
    fmt.Println(animal.Speak())

    //animal = Weird{"The Human"}
    //fmt.Println(animal.Speak())
}