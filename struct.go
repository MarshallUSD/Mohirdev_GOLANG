package main

import "fmt"

type Person struct{
	Name string
	Age int
	city string
	is_alive bool
}

func (p Person) displayInfo(){
	fmt.Printf("Name: %s\nAge: %d\nCity: %s\nIs Alive: %t", p.Name, p.Age, p.city, p.is_alive )

}


func main(){
	persona := Person{
		Name: "John Doe",
		Age: 30,
		city: "New York",
		is_alive: true,
	}
	persona.displayInfo()

}

