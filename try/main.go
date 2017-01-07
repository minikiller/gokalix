package main

import "fmt"

func main() {
	a,b:=5,6
	//f:=sum
	function(a,b,sum)
}

func function (x,y int, f func (int,int) int){
	fmt.Println(f(x,y))
}

func sum (x,y int) int {
	return x+y
}
