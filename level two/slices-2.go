package main

import (
	"fmt"
	"strings"
	"sort"
)


func main(){

	names := strings.Split("venkatesh vinodh venki vinod kumar"," ");
	
	var slice1 []string;
	
	slice1 = append(slice1,names...);
	
	sort.Strings(slice1);
	
	fmt.Println(slice1);
	
	sort.Strings(names);
	
	fmt.Println(names);	

	slice2 := append([]string(nil),slice1...)
	
	fmt.Println(slice2);
	
	slice3 := make([]string,10);
	
	fmt.Printf("%#v\n",slice3);
	
	slice3 = append(slice3,slice3...);
	
	fmt.Printf("%q\n",slice3);
	
	var slice4 = []int{11,22,33};
	
	var slice5 = []int{1,2,3};
	
	slice := copy(slice4,slice5);
	
	fmt.Println("Number of elements copied",slice);
	fmt.Println(slice4);

}
