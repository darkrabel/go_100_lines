package main
import "fmt" 
var a int
//   declaring  a intiger  varriable.
 var(
	 i int
	 b bool
	 c float32
	 s string
 )


func main() {
	a = 42               /* assign single value*/
	b,c = true,32.4 /* assign multiple value in single line*/
    s = "go " /* assigning string  and string must contain double cotation*/
	fmt.Println(a,b,c,s)
}