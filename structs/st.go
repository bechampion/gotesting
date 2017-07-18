package main

import "fmt"

// Nick() becomes a func within a struct
func (p *Person) Nick() string {
	p.Age = 20
	return "Manio"
}

type Person struct {
	Name string
	Age  int
}

type Android struct {
	Person
	Model string
}

func main() {
	fmt.Println("vim-go")
	jeronimo := Person{Name: "jeronimo", Age: 19}
	fmt.Println(jeronimo.Age)
	//// struct with a func , like a class
	fmt.Println(jeronimo.Nick())
	android := Android{Person{Name: "Robot", Age: 2000}, "v11111"}
	fmt.Println(android.Model)
}
