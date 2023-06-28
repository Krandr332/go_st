package main

import "fmt"
// указатели
func main(){
	a:= 1
	b:= &a
	*b +=1
	fmt.Println(a,*b)}