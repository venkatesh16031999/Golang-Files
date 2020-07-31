package main

import (
	"fmt"
	"strings"
	"math/rand"
	"time"
)

func main(){

	rand.Seed(time.Now().UnixNano());

	var words []string = strings.Fields("hello my name is venkatesh rajendran");
	
	fmt.Println(words);
	
	for i := 0 ; i < len(words) ; i++ {
		fmt.Println(words[i]);
	}
	
	for _,v := range words{
		fmt.Println(v);
	}
	
	for _,v := range words[len(words)-1:] {
		fmt.Println(v);
	}
	
	fmt.Println(rand.Intn(10));
	
	for _,v := range words[1:] {
		fmt.Println(v);
	}

}
