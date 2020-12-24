package main
 
import (
    "fmt"
    "errors"
)
 
func e(v int) (int, error) {
    return 42, errors.New("42 is unexpected!")
}
 
func main() {
    _, err := e(0)
     
    if err != nil {   // check error here
        fmt.Println(err)      // 42 is unexpected!
    }   
}
