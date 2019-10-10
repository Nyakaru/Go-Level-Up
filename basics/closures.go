package main
import "fmt"

func add() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
//Functional prograammming
func main() {
	sum :=  add()
	for i := 0; i <= 10; i++ {
		fmt.Println(sum(i))
	}
}
