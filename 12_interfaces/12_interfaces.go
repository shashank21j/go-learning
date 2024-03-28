package main

type Name interface {
	GetName() string
}

type Person struct {
	firstName string
	lastName  string
}

func (p Person) GetName() string {
	return p.firstName + " " + p.lastName
}

func (p Person) SayHi() {
	println("Hi, I am " + p.GetName())
}

type Employee struct {
	name string
}

func (e Employee) GetName() string {
	return e.name
}

func PrintName(obj Name) {
	println(obj.GetName())
}

func main() {
	var name Name = Person{"John", "Doe"}
	println(name.GetName())

	names := []Name{Person{"Alice", "Doe"}, Employee{"Bob"}}
	for _, n := range names {
		println(n.GetName())
	}

	var p1 Name = Person{"James", "Doe"}
	PrintName(p1)
	Person{"Alice", "Doe"}.SayHi()
}
