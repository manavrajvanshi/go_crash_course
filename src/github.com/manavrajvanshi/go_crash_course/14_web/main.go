package main

import (
	"fmt"
	"net/http"
)

func index( w http.ResponseWriter, r *http.Request){
	fmt.Fprintf( w, "Hello World from GoLang server")
	fmt.Println( r )
}


func main(){
	http.HandleFunc( "/", index)
	fmt.Println("Server starting")
	http.ListenAndServe( ":3005",nil)
}
