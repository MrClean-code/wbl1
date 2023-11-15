package main

import "fmt"

type Human struct {
	ID      int8
	Name    string
	Surname string
}

func (h *Human) live() {
	fmt.Printf("Я %s %s", h.Name, h.Surname)
}

type Action struct {
	Name string
	Human
}

func NewAction(name string, human Human) *Action {
	return &Action{Name: name, Human: human}
}

func (a *Action) move() {
	fmt.Println("Я бегаю")
}

func main() {
	run := NewAction(
		"спорт",
		Human{
			ID:      1,
			Name:    "Михаил",
			Surname: "Сорокин",
		})

	run.move()
	run.live()
}
