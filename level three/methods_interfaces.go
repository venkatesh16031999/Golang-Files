package main 

import (
	"fmt"
)

type book struct{
	name string
	price int
}

type pricer interface{
	printPrice();
}

type list []pricer

func main(){

	book1 := book{
		name: "book1",
		price: 100,
	}
	
	book2 := book{
		name:"book2",
		price:20,
	}
	
	book1.printPrice();
	book.printPrice(book1);
	
	store := list{}
	
	store = append(store,book1,book2);
	
	fmt.Println(store);
	
	store.printPrice();
	
}


func (l list) printPrice(){
	for _,v := range l{
		v.printPrice();
		
		b,isBook:=v.(book)
		
		if(isBook){
			fmt.Println(b);
		}
		
		switch v:= v.(type) {
			default:
				fmt.Println(v);
		}
		
	}

}

func (b book) printPrice(){
	fmt.Println("the price of the book",b.price);
}
