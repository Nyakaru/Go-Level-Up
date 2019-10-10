package main
import "fmt"

func main() {

	nums := [] int {23,43,54,26,76,87,98}

	for i, num := range nums {
		fmt.Printf("%d -id: %d\n", i, num)
	}

	//Not using index
	for _, num := range nums {
		fmt.Printf("id: %d\n", num)
	}

	//Add ids together
	sum := 0
	for _, num := range nums {
		sum += num
}
fmt.Println("Sum", sum)

//Using maps

staff := map[string]string {"Chef":"Ken","Tutor":"Kev"}

for k, v := range staff {
	fmt.Printf("%s : %s\n", k, v)
}
for k := range staff {
	fmt.Println("Title:" + k)
}
}
