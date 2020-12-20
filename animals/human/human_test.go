package human

import (
	"testing"
	"regexp"
)

//TestSpeak calls human.Speak with a name,
//checking for a valid return value
func TestSpeak(t *testing.T) {
	name := "Hello human"
	animal := Human{name}
	want := regexp.MustCompile(`\bI am a human, my name is `+name+`\b`)
	msg, err := animal.Speak()
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Speak() = %q, %v, want match for  %#q, nil`, msg, err, want)
	}
}

//TestSpeakError calls human.Speak with and empty 
//string checking for an error.
func TestSpeakError(t *testing.T) {
	name := ""
	animal := Human{name}
	msg, err := animal.Speak()
	if msg != "" || err == nil {
		t.Fatalf(`Speak() = %q, %v, want "", error`, msg, err) 
	}
}