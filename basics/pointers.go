package main
import "fmt"

func main() {
	x := 7
	y := &x

	fmt.Println(x,y)
	fmt.Printf("%T\n",y)

	//Use * to read the values
	fmt.Println(*y)
	fmt.Println(*&x)

	//Change the value with pointer
	*y = 5
	fmt.Println(x)
}
