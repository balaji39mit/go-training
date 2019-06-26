package oops

//Run-time polymorphism is achieved through interfaces.
//Compile-time polymorphism is not supported in Go to make the code more non-readable.

import "fmt"

type Speaker interface {
	Speak() string
}

type Human struct {
	sound string
}

type Dog struct {
	sound string
}

type Cat struct {
	sound string
}

type Alien struct {
	sound string
}

func (h Human) Speak() string {
	return h.sound
}

func (d Dog) Speak() string {
	return d.sound
}

func (c Cat) Speak() string {
	return c.sound
}

func (a Alien) Speak() string {
	return a.sound
}

func getType(i interface{}) string {
	switch i.(type) {
	case Dog:
		return "Dog"
	case Cat:
		return "Cat"
	case Alien:
		return "Alien"
	case Human:
		return "Human"
	default:
		return "Unknown"
	}
}

func init() {
	h := Human{sound: "Hello"}
	d := Dog{sound: "bow bow"}
	c := Cat{sound: "meow meow"}
	a := Alien{sound: "weird sound"}
	arr := []Speaker{h, d, c, a}
	for _, value := range arr {
		fmt.Println(fmt.Sprintf("Type %v sounds like %v", getType(value), value.Speak()))
	}
}
