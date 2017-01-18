package main

import (
	"fmt"
)

type Developers []Developer

type Developer interface {
	Display() string
}

type Java struct {
	Name string
	Age  int
	Says string
}

func (j Java) Display() string {
	return j.Name + " Says " + j.Says
}

type Go struct {
	Name string
	Age  int
	Says string
}

func (g Go) Display() string {
	return g.Name + " Says " + g.Says
}

type Php struct {
	Name   string
	Age    int
	Shouts string
}

func (p Php) Display() string {
	return p.Name + " Shouts " + p.Shouts
}


type C struct {
	Name     string
	Age      int
	Explains string
}

func (c C) Display() string {
	return c.Name + " Explains " + c.Explains
}


func main() {
	p := Developers{
		Php{Name: "connor", Age: 22, Shouts: "I like things slow and loosely typed"},
		Go{Name: "charlie", Age: 24, Says: "Throw me a Gopher"},
		Java{Name: "alex", Age: 25, Says: "I have a factory full of problems"},
		Go{Name: "bill", Age: 55, Says: "Dont talk to me about dependency management"},
		C{Name: "Jill", Age: 45, Explains: "I wish my life was simpler"},
		C{Name: "Ahamed", Age: 33, Explains: "Im tradtional and so is my language"},
		Java{Name: "Rob", Age: 55, Says: "Your only using up 99% of your resources..."},
		Php{Name: "Kimbley", Age: 23, Shouts: "Tabs and spaces lets spend loads of time debating it..."},
	}

	for _,k := range p {
		fmt.Println(k.Display())
	}
}
