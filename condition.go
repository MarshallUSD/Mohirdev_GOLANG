package main

import "fmt"


func main(){
/*
	x:=0

	if x>0{
		fmt.Println("X is positive")
	 } else if x==0{
		fmt.Println("X is zero")
	 }else{
		fmt.Println("X is negative")
	 }
*/

// switch in go
 /*d := 5
 switch d {
 case 1:
	    fmt.Println("Monday")
 
case 2:
	fmt.Println("Tuesday")

	case 3:
	fmt.Println("Wednesday")
case 4:
	fmt.Println("Thursday")
case 5:
	fmt.Println("Friday")
case 6:
	fmt.Println("Saturday")
case 7:
	fmt.Println("Sunday")
default:
	fmt.Println("Invalid day")	
}
	*/

   n:= 0

   switch{
   case n%2==0:
		fmt.Println("Even number")
   case n%2!=0:
		fmt.Println("Odd number")
	case n==0:
		fmt.Println("Number is zero")
    default:
		fmt.Println("Invalid number") 
	}
}


