package main

import "fmt"

func main(){
	a := 5

	b := &a


	fmt.Println( a, b)
	fmt.Printf( "%T\n%T\n",a , b)

	fmt.Println( *b )

	fmt.Println( b)

}
