package supertype

type SpeakBehaviour interface {
    Speak() (string, error) 
}