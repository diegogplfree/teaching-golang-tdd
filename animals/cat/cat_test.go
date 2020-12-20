package cat

import (
	"testing"
	"regexp"
)

//TestSpeak calls cat.Speak with a name,
//checking for a valid return value
func TestSpeak(t *testing.T) {
	name := "Meow cat"
	animal := Cat{name}
	want := regexp.MustCompile(`\bI am a cat, my name is `+name+`\b`)
	msg, err := animal.Speak()
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Speak() = %q, %v, want match for  %#q, nil`, msg, err, want)
	}
}

//TestSpeakError calls cat.Speak with and empty 
//string checking for an error.
func TestSpeakError(t *testing.T) {
	name := ""
	animal := Cat{name}
	msg, err := animal.Speak()
	if msg != "" || err == nil {
		t.Fatalf(`Speak() = %q, %v, want "", error`, msg, err) 
	}
}