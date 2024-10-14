package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) Talk() {
	fmt.Printf("Hello, my name is %s and i'm %d y.o.\n", h.Name, h.Age)
}

type Action struct {
	Human
	Action string
}

func (a *Action) doAction() {
	fmt.Printf("I'm %s\n", a.Action)
}

func main() {
	human := Action{
		Human: Human{
			Name: "Maks",
			Age:  22,
		},
		Action: "walking",
	}

	human.Talk()
	human.doAction()
}
