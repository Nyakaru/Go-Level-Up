package main
import "fmt"

func main() {
	// Main types
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for uint32
	// float32 float64
	// complex64 complex128

	// using Var

	var name = "Level"
	var version = 1
	var length float32 = 8.9
	var isRead = false
	isRead = true

	//shorthand 
	value := "scratch"
	size := 1.89

	indetity, email := "kinara" , "kin@gmail.com"

	fmt.Println(name, version, isRead, value, size, length, indetity, email)
	fmt.Printf("%T\n",name)
	fmt.Printf("%T\n",version)
	fmt.Printf("%T\n",isRead)
	fmt.Printf("%T\n",value)
	fmt.Printf("%T\n",size)
	fmt.Printf("%T\n",length)
}
