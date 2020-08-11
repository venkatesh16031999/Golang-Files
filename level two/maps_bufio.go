package main

import (
    "fmt"
    "os"
    "strings"
    "bufio"
    )
    
func main(){
    
    io := bufio.NewScanner(os.Stdin);
    
    io.Scan();
    
    fmt.Println(strings.ToUpper(io.Text()));
    
    maps := map[string]bool{}
    
    maps["hello"]=true;
    
    fmt.Println(maps["hello"]);
    
    var map1 = map[string]string{};

    var n int = 2;
    
    for i:= 0 ; i<n ; i++ {
        io.Scan();
        
        map1[io.Text()]=io.Text();
        
    }
    
    delete(map1,io.Text())
    
    var slice1 = []int{1,2,3};
    
    fmt.Println(slice1);
    
    fmt.Println(map1);
    
    os.Stdin.Close()
    
}
