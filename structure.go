package main

import "fmt"

// Student structure
type Student struct {
    firstName, lastName string
    age                 int
}

func main() {

    // creating first student structure using field names
    stud := Student{
        firstName: "Krunal",
        lastName:  "Lathiya",
        age:       26,
    }

    //creating second structure without using field names
    stud2 := Student{"Ankit", "Lathiya", 25}

    fmt.Println("Student 1", stud)
    fmt.Println("Student 2", stud2)
}
