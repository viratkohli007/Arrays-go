package main
import "fmt"

func main()  {

	var a [5]int
	fmt.Println(a)
	a[4] = 100
	fmt.Println("set a: ",a[4])
	fmt.Println("get Array: ", a)
	fmt.Println("length: ",len(a))

	b := [5]int{1,2,3,4,5}
	fmt.Println("b is: ", b)

	var twoD [2][3] int 
	for i:=0; i<2; i++{
		for j:=0; j<3; j++{
			twoD[i][j] = i+j
			
		}
	}
	fmt.Println(twoD)
}