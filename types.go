package main
import "fmt" 

func main() {
  var c complex128 = 1 + 4i  // 128 bit complex number.
  var i int32 = 4 // 32 bit intiger
  var f float32 = 1.0 // 32 bit floeat
  var d uint =  3 // 16 bit unsigned intiger.
  /* Default types */
  n := 42              // int
  pi := 3.14           // float64
  x, y := true, false  // bool
  z := "Go is awesome" // string
 fmt .Println(c,i,f,d)
 fmt.Println(n,pi,x,y,z)
}