package main
import "fmt"
/*
func printme(){
	fmt.Println("Hello from printme function")
}

func main(){
	printme()
}
	*/

func describe_me(name string, age int){
	fmt.Println("my name is ", name, "and my age is", age)
}	

func add(a int, b int ) int {
	return a+b
}

func main(){
    fmt.Println(add(5,4))
	describe_me("Alim", 19)
}





