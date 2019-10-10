package main
import "fmt"

func main() {
	// Arrays
	var staffArr [3]string

	// Assign values
	staffArr[0] = "Peter"
	staffArr[1] = "Parker"
	staffArr[2] = "Baker"
	
	//Declare and assign
	animalArr := [3]string{"Spider","Ant","Ewe"}

	vowArr := []string{"a","b","c","d","e"}

	fmt.Println(animalArr,staffArr[1:2],len(vowArr))
}
