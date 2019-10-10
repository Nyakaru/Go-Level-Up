package main
 import "fmt"

 func main() {
	 //longer method
	 i :=1

	 for i <= 10 {
		 fmt.Println(i)
		 i++
	 }

	 // shorter method
	 for i :=1; i <= 10; i++  {
		 fmt.Printf("Num %d \n", i)
	 }

	 //Fizzbuzz
	 for i :=1; i <= 100; i++ {
		 if i % 3 == 0 {
			 fmt.Printf("Fizz \n")
		 } else if i % 5 == 0 {
			fmt.Printf("Buzz \n")
		 } else if i % 7 == 0 {
			fmt.Printf("FizzBuzz \n")
		 } else {
			fmt.Printf("%d \n", i)
		 }
	 }
 }
