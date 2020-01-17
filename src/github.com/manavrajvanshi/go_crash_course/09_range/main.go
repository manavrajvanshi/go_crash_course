package main

import "fmt"

func main(){
	ids := []int{33,44,55,66,77,88,34,635,6342,4,6264563,564}
	
	
	for _,id:= range ids {
		fmt.Printf( "ID: %d\n", id)
	}

	sum := 0

	for _, id := range ids{
		sum += id
	}

	fmt.Printf( "SUM of IDS: %d\n", sum)
	

	emails := map[string]string { "Alice":"a@a.com", "Bob":"b@b.com", "Charles":"c@c.com" } 

	for k,v := range emails{
		fmt.Printf( "Name: %s, Email: %s\n", k, v)
	}

	for k := range emails {
		fmt.Printf( "Name: %s, Email: %s\n", k, emails[k])
	}
}
