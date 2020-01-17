package main

import "fmt"

func main (){

	for i:= 0; i <=100; i++ {
		if (i%5==0) && (i%3 == 0){
			fmt.Printf( "%d is FizzBuzz\n", i)
		}else if i%5==0{
			fmt.Printf( "%d is Buzz\n" , i)
		}else if i%3==0{
			fmt.Printf( "%d is Fizz\n", i)
		}else{
			fmt.Println(i)
		}
	}
}
