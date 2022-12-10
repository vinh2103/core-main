package main

import "fmt"

// this is entity
type person struct {
	age  int
	name string
}

// this is data of entities
type data struct {
	people []person
}

// this is interface of data controller
type dataController interface {
	add(person) person
	get() []person
}

func (d *data) add(p person) person {
	d.people = append(d.people, p)
	return p
}
func (d *data) get() []person {
	return d.people
}

func main() {
	a := person{
		age:  182,
		name: "vinj",
	}
	b := &data{}
	b.add(a)

	fmt.Print(b.get())

}
