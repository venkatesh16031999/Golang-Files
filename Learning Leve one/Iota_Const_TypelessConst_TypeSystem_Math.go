package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	const (
		incValues1 = (iota+1) * 10
		incValues2
		incValues3
		incValues4
		incValues5
	)
	
	fmt.Println(math.Pow10(incValues5));
	
	var floatvar float64 = 1.6132;
	
	fmt.Println(math.Round(floatvar));
	fmt.Println(math.Floor(floatvar));
	fmt.Println(math.Ceil(floatvar));
	fmt.Println(math.Sqrt(floatvar));
	
	type definedInt int64
	
	var number definedInt = 10;
	
	fmt.Println(number);
	fmt.Println(time.Now());
}
