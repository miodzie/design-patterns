package main

import (
	"fmt"
	"strings"
)

type Command interface {
	Execute()
}

type MenuItem struct {
	DisplayName string
	Command
}

func (m MenuItem) String() string {
	return m.DisplayName
}

type Menu struct {
	Items []MenuItem
}

func (m *Menu) Add(item MenuItem) *Menu {
	m.Items = append(m.Items, item)
	return m
}

func (m Menu) String() string {
	b := strings.Builder{}
	for i, mi := range m.Items {
		b.WriteString(fmt.Sprintf("%d: %s\n", i+1, mi.String()))
	}
	return b.String()
}

func (m *Menu) HasItemIndex(i int) bool {
	return i <= len(m.Items)
}

type GreetCommand struct{}

func (g GreetCommand) Execute() {
	var name string
	fmt.Println("What is your name?")
	fmt.Scan(&name)
	fmt.Printf("Hello %s!", name)
}

func main() {
	menu := Menu{}
	menu.Add(MenuItem{DisplayName: "Hello World!", Command: HelloWorldCommand{}}).
		Add(MenuItem{DisplayName: "Hello World!", Command: HelloWorldCommand{}}).
		Add(MenuItem{DisplayName: "Greeter", Command: GreetCommand{}})

	fmt.Print(menu)

	var selectedItem int
	fmt.Scan(&selectedItem)
	if selectedItem != 0 && menu.HasItemIndex(selectedItem) {
		menu.Items[selectedItem-1].Execute()
	}
}

type HelloWorldCommand struct {
}

func (p HelloWorldCommand) Execute() {
	fmt.Println("Hello world!")
}
