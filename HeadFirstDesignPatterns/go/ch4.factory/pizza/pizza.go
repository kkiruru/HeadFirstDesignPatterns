package pizza

import (
	"fmt"
)

type making interface {
	Prepare()
}

type Pizza struct {
	Name string
	Dough string
	Sause string
	making
}


func (p *Pizza) Bake() {
	fmt.Println("Bake for 25 minutes at 350")
}


func (p *Pizza) Cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}


func (p *Pizza) Box() {
	fmt.Println("Place pizza in official PizzaStore box")
}


type NYStyleCheesePizza struct {
	Pizza
}


func (n *NYStyleCheesePizza) Prepare() {
	fmt.Printf("Preparing: %s \n", n);
	fmt.Println("Tossing dough... ");
}


func init() {
	fmt.Println("init of pizza package");
}


func (n *NYStyleCheesePizza) init() {
	fmt.Println("init of NYStyleCheesePizza struct");
	n.Pizza = Pizza {
		Name: "New York Style Cheese Pizza",
		Dough: "Thin Crush",
		Sause: "Marinara",
		}
}
