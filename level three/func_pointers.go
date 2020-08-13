package main 

import (
	"fmt"
);

func main(){
	
	subFunc();
	sub1(1);
	value,error := sub2(2);
	fmt.Println(value,error);
	value1,error1 := sub3();
	fmt.Println(value1,error1);
	
	n := 10;
	ptr(&n);
	fmt.Println(n);
}

func ptr(n *int){
	(*n)++;
	fmt.Printf("%p",n);
}

func sub1(a int) int {
	fmt.Println(a);
	return a;
}

func subFunc(){
	fmt.Println("calling func");
}

func sub2(a int) (int, error){
	fmt.Println(a);
	er := fmt.Errorf("error is printed");
	return a,er;
}

func sub3() (a int,e error){
	e = fmt.Errorf("error is printed2");
	a=1;
	return 
}
