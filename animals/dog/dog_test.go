package dog

import (
	"testing"
	"regexp"
)

//TestSpeak calls dog.Speak with a name,
//checking for a valid return value
func TestSpeak(t *testing.T) {
	name := "Bark dog"
	animal := Dog{name}
	want := regexp.MustCompile(`\bI am a dog, my name is `+name+`\b`)
	msg, err := animal.Speak()
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Speak() = %q, %v, want match for  %#q, nil`, msg, err, want)
	}
}

//TestSpeakError calls dog.Speak with and empty 
//string checking for an error.
func TestSpeakError(t *testing.T) {
	name := ""
	animal := Dog{name}
	msg, err := animal.Speak()
	if msg != "" || err == nil {
		t.Fatalf(`Speak() = %q, %v, want "", error`, msg, err) 
	}
}