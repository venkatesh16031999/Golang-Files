package main

import (
	"fmt"
)

func main(){

	var arr [3]int;
	
	arr[0]=1;
	arr[1]=2;
	arr[2]=3


	fmt.Println(arr);
	
	var arr2 = [...]string{
	"1",
	"2",
	"3",
	}
	
	fmt.Println(arr2);
	
	var arr3 = [3]string{
	1:"1",
	0:"2",
	2:"3",
	}
	
	fmt.Println(arr3);
	
	type arrtype [3]string;
	
	var ar = arrtype{"1","2","3"}
	
	fmt.Println(ar);
	fmt.Println(ar==[3]string{"1","2","3"})
	
	var mul = [3][3]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	
	fmt.Println(mul);

}
