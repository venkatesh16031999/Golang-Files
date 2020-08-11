package main;

import (
	"fmt"
	"unicode/utf8"
)

func main(){

	var str string = "hello";
	
	for _,v := range str {
		fmt.Print(string(v));
	}
	fmt.Println();

	value,size := utf8.DecodeRune([]byte(str));
	
	fmt.Println(value," ",size);
	
	var bytes []byte = []byte(str);
	
	fmt.Println(bytes);
	
	var strings string = string(bytes);
	
	fmt.Println(strings);
	
	buf := make([]byte,0,len(str));
	
	for i,v := range str {
		buf = append(buf,str[i]);
		fmt.Printf("%[1]T : %[1]v  ;  %[2]T : %[2]v\n",str[i],v);
	}
	
	fmt.Println(buf);
	
}
