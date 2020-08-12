package main 

import (
	"fmt"
	"encoding/json"
)

type globalStruct struct {
	year int
	typeName string
	cost int
}

func main(){

	type book struct{
		globalStruct
		author string
		typeName string
	}
	
	book1 := book{
		globalStruct : globalStruct{year : 1990,typeName : "comics",cost : 200},
		author : "venkatesh",
		typeName: "parent type name",
	}
	
	fmt.Println("globalStruct name",book1.globalStruct.typeName);
	
	fmt.Println("parent name",book1.typeName);
	
	fmt.Println(book1);

	type jsonStruct struct{
		Name string
		Age int
		Experiance int 
		Money int
	}
	
	value := jsonStruct{
		Name: "Venkatesh.R",
		Age: 21,
		Experiance : 1,
		Money : 15000,
	}
	
	fmt.Println(value);
	
	jsonValue, err := json.Marshal(value);
	
	if err != nil {
		fmt.Println("invalid json type");
		return;
	}
	
	fmt.Println(string(jsonValue));

}
