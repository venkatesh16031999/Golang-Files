package main

import (
	"fmt"
	"time"
	"strings"
	"strconv"
)

func main(){

	floatVal := "1.2123";

	var err error;

	fmt.Println(time.Now().Hour());
	fmt.Println(strings.TrimSpace("  Hello  "));
	floatString,err:=strconv.ParseFloat(floatVal,64);
	
	if(err!=nil){
		fmt.Println("The Value is in correct");
	}else{
		fmt.Println(floatString);
	}
	
	if floatString==1.2123 {
		fmt.Println("it is converted from string to the float");
	}
	else if floatString != 1.2123 {
		fmt.Println("It is still remaining as a string");
	}else{
		fmt.Println("Nothing Happened");
	}
}
