package main 
  
import ( 
    f "fmt"
    "regexp"
) 
  
func main() { 
    f.Println("--------Reference strings--------\n") 
  
    name := "My name is Geeks for geeks."
    f.Println(name) 
  
    profession := "I am a computer science portal for geeks."
    f.Println(profession) 
  
    message := "You can find anything here, if not tell us and we'll add it for you!"
    f.Println(message) 
  
    //--------------------------------------------------- 
    f.Println("\n--------Match functions--------") 
    //------------------------------------------- 
  
    obj, err := regexp.Match("[gG]e*k.*", []byte(name)) 
    f.Println("\nregex.Match returns ->", obj, 
        "and error(if any) --->", err) 
  
    //------------------------------------------- 
    obj, err = regexp.Match("[gG]e*k.*", []byte(profession)) 
    f.Println("\nregex.Match returns ->", obj, 
        "and error(if any) --->", err) 
  
    //------------------------------------------- 
    obj, err = regexp.MatchString("Geek.*", message) 
    f.Println("\nregex.MatchString returns ->", obj, 
        "and error(if any) --->", err) 
  
    //------------------------------------------- 
}
