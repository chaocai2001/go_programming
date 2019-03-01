package composite_inheritance

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	Pet
}

func (p *Dog) Speak() {
	fmt.Print("...")
}

func TestSubClassAccess(t *testing.T) {
	aDog := new(Dog)
	aDog.SpeakTo("Chao")
}

func makePetSpeak(p Pet) {
	p.Speak()
}

func TestLSP(t *testing.T) {
	aDog := new(Dog)
	//makePetSpeak(aDog) //cannot use aDog (type *Dog) as type Pet
	//makePetSpeak(Pet(aDog)) // cannot convert aDog (type *Dog) to type Pet
	makePetSpeak(aDog.Pet)
}
