package main

import "fmt"

type details struct {
	age     int
	dept    string
	floor   int
	manager string
}

var People = map[string]details{
	"charlie": details{
		age:     22,
		dept:    "development",
		manager: "Fraser",
		floor:   5,
	},
	"john": details{
		age:     34,
		dept:    "Product",
		manager: "Chris",
		floor:   5,
	},
	"bob": details{
		age:     24,
		dept:    "Development",
		manager: "Charlie",
		floor:   5,
	},
	"chris": details{
		age:     34,
		dept:    "Development",
		manager: "Fraser",
		floor:   3,
	},
	"nadim": details{
		age:     44,
		dept:    "Development",
		manager: "Fraser",
		floor:   3,
	},
	"josh": details{
		age:     25,
		dept:    "Development",
		manager: "Jai",
		floor:   5,
	},
	"micheal": details{
		age:     28,
		dept:    "Development",
		manager: "Charlie",
		floor:   5,
	},
	"alex": details{
		age:     32,
		dept:    "Development",
		manager: "Charlie",
		floor:   5,
	},
	"andrew": details{
		age:     44,
		dept:    "Development",
		manager: "Charlie",
		floor:   3,
	},
}

func function1() []string {
	devIn5th := make([]string, 0)
	for v, k := range People {
		if k.dept == "Product" && k.floor == 5 {
			devIn5th = append(devIn5th, v)
		}
	}
	return devIn5th
}

func function2() []string {
	charlieTeam := make([]string, 0)
	for v, k := range People {
		if k.manager == "Charlie" {
			charlieTeam = append(charlieTeam, v)
		}
	}
	return charlieTeam
}

func function3() []string {
	productTeam := make([]string, 0)
	for v, k := range People {
		if k.dept == "Product" {
			productTeam = append(productTeam, v)
		}
	}
	return productTeam
}

func main() {

	fmt.Println("Developers in 5th floor:")
	fmt.Println(function1())
	fmt.Println("People in Charlie's Team:")
	fmt.Println(function2())
	fmt.Println("People in Product Dept:")
	fmt.Println(function3())
}
