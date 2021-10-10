// Notice how each variable declaration is followed by the type of that variable. Before we cover types in the next section, note that we can replace the var keyword with const when
// we need to introduce constants in our code.
// When declaring variables, another option is to use the := operator to initialize and assign to variables in one go. This is called a short variable declaration. Let’s refactor the code above to illustrate thi
package main

import "fmt"

func main() {
	i := 4             // intailize and assign value.
	b, c := true, 32.0 // intialize and assign multiple value.
	s := "Go is awesome"
	fmt.Println(i, b, c, s)
}
