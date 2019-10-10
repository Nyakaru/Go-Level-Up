package main
import ("fmt"
        "strconv")

// Define employee struct
type Employee struct {
	firstName, lastName, city, gender string
	age int
}

//Greeting method(value receiver)
func (person Employee) greet() string {
	return "Hello my name is " + person.firstName + " " + person.lastName + " and I am " + strconv.Itoa(person.age)
}

//Birthday method(pointer receiver)
func (person *Employee) birthday() {
	person.age++
}

//Married method(pointer receiver)
func (person *Employee) married(spouseName string) {
	if person.gender == "M" {
		return
	} else {
		person.lastName = spouseName
	}
}

func main() {
	
	//Init employee
	emp1 := Employee{firstName:"Kin",lastName:"Nyaks",city:"Nai",gender:"M",age:25}

	//Alternative
	emp2 := Employee{"Agg","Kas","Mos","F",24}
	
	fmt.Println(emp1,emp2)
	
	emp1.age++
	fmt.Println(emp1.age)

	emp1.birthday()
	fmt.Println(emp1.greet())

	emp2.married("Keps")
	fmt.Println(emp2.greet())
}
