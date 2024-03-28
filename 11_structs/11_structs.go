package main

import "fmt"

/*
Structs are typed collections of fields. They’re useful for grouping data together to form records.
Capitalized fields are exported, meaning they can be accessed from other packages.
*/
type Person struct {
	name string
	age  int
}

// class method
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

// Inheritance example
type Doctor struct {
	Person
	speciality string
}

// class method
func (d Doctor) String() string {
	return fmt.Sprintf("%v, %v", d.Person, d.speciality)
}

func (p Person) Equals(p2 Person) bool {
	return p.name == p2.name && p.age == p2.age
}

func main() {
	p1 := Person{"James", 25}
	fmt.Println(p1)
	fmt.Println(p1.name)
	fmt.Println(p1.age)

	d1 := Doctor{
		Person:     Person{"John", 35},
		speciality: "Cardiology",
	}
	fmt.Println(d1)

	p := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	fmt.Println(p[0].Equals(p[0]))
}
