package main

import (
	"strings"
	"fmt"
	"unicode/utf8"
	"os"
)

func main(){

	var a,b int = 10,20;
	externalCall(a,b);
	var myName string = "Venkatesh";
	
	var args1 []string = os.Args
	
	args2 := os.Args
	
	fmt.Println("arguments length of 1 :",len(args1));
	
	fmt.Println("arguments length of 2 :",len(args2));
	
	fmt.Println("My Name Is",myName);
	fmt.Println("My Nmae Length",utf8.RuneCountInString(myName));
	fmt.Println("Converted to Upper Case",strings.ToUpper(myName));
	fmt.Println("Converted to Lower Case",strings.ToLower(myName));
	
}

func externalCall(a int, b int){

	fmt.Println("externalFunction called with paramters",a,b);

}
