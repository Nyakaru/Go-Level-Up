package main

import "fmt"

func main() {
	x := 5
	y := 7

	// If else
	if x > y {
		fmt.Printf("%d is greater than %d \n",x,y)
	} else {
		fmt.Printf("%d is greater than %d \n",y,x)
	}

	//else if
	complexion := "black"

	if complexion == "red" {
		fmt.Println("it is red")
	} else if complexion == "blue" {
		fmt.Println("it is blue")
	} else {
		fmt.Println("neither blue nor red")
	}

	//switch
	// &&-and ||-or 
	switch complexion {
	case "red":
		fmt.Println("it is red")
	case "blue":
		fmt.Println("it is blue")
	default:
		fmt.Println("neither blue nor red")
	}
}
