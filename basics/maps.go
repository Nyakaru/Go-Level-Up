package main
import "fmt"

func main() {
	// Define map
	emails := make(map[string] string)

	//Assign key values
	emails["kin"] = "kin@gmail.com"
	emails["nyaks"] = "nyaks@gmail.com"

	fmt.Println(emails, emails["nyaks"])

	//Delete from map
	delete(emails, "nyaks")
	fmt.Println(emails, len(emails))

	//Declare map and add kv
	staff := map[string]string {"Chef":"Ken","Tutor":"Kev"}
	fmt.Println(staff)

}
