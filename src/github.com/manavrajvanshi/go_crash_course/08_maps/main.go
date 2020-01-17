package main

import "fmt"

func main(){
	emails := make( map[string] string)
	emails["Bob"] = "bob@gmail.com"
	emails["Charles"] = "charles@gmail.com"
	emails["Alice"] = "alice@gmail.com"
	fmt.Println(emails)
	fmt.Println("Deleting Alice")
	delete( emails, "Alice")
	fmt.Println( emails )
	fmt.Println( emails["Alice"])
	fmt.Println("-------------------")

	names := map[string]string { "Manav":"Rajvanshi", "Mehul":"Rajvanshi"}
	fmt.Println( names)
}
