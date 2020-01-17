package main

import "fmt"


func greeting ( name string) string {
	return "Hello " + name
}

func getSum ( x , y int) (sum int){
	sum = x + y
	return
}

func main(){
	fmt.Println(greeting( "Manav" ) )
	fmt.Println( getSum( 324123, 123412341) )
}
