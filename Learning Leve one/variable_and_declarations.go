package main

import (
	"fmt"
	"path"
	timeLib "time"
	"strconv"
)

// multiple variable declaration
var (
	name string = "Venkatesh"
	age int =  21
	place string = "India"
)

func main(){
	// short decalaration
	a:=12;
	b:=21;
	
	speed,time := 10,20;
	
	urlPath := "css/index.css"
	
	// _ --> blank identifier
	 
	_,fileName := path.Split(urlPath);
	
	currentTime := timeLib.Now();
	
	fmt.Println("Time :", currentTime);
	
	fmt.Println("File Name :",fileName);
	
	fmt.Println("Speed :",speed,"Time :",time);
	
	fmt.Println(a,b);
	printDetails();
	
	num := 20;
	
	fmt.Println(strconv.Itoa(num));
	
}

func printDetails(){
	fmt.Println("Name: ",name,"Age: ",age,"country: ",place);
}
