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

func main() {

	fmt.Println("Employees working in Development in 5th floor:")
	for v, k := range People {
		if k.dept == "Product" && k.floor == 5 {
			fmt.Println(v)
		}
	}

	fmt.Println("Employees having Charlie as manager:")
	for v, k := range People {
		if k.manager == "Charlie" {
			fmt.Println(v)
		}
	}

	fmt.Println("Employees working on Product:")
	for v, k := range People {
		if k.dept == "Product" {
			fmt.Println(v)
		}
	}

}
