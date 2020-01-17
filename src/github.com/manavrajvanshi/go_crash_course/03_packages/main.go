package main

import (
	"fmt"
	"math"
	"github.com/manavrajvanshi/go_crash_course/03_packages/strutil"
)


func main(){
	fmt.Println( "Hello, world" )
	fmt.Println( math.Floor( 2.432424) )
	fmt.Println( math.Ceil( 32423.3423423) )
	fmt.Println( math.Sqrt( 26 ) )
	fmt.Println( strutil.Reverse("!olleh" ))
}
