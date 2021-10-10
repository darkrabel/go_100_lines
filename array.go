package main
import "fmt" 

func main() {
	//using array.
var devolopmentoption = [4]string {"r-pi" , "aws ", "azure ","google_cloud"}
//  loop thourgh aray 
 for I := 0; I < len(devolopmentoption); I++ {
	 option := devolopmentoption[I]
	 fmt.Println(option)
	 
 }
}