package main

import "fmt"

// Заполнение массива от от 1 до s
func z1()([]int,int){
	var s int
	var slice []int
	// var number int
	fmt.Println("Введите размер массива")
	_, err := fmt.Scanln(&s)
	if err != nil {
		fmt.Println("Ошибка")
		return slice, s
	}
	fmt.Println(s)
	for i := 0; i != s+1; i++ {
		slice = append(slice, i)
		fmt.Println(slice)
	}
	return slice, s
}
// сортировка по четности
func z2() {
	var slice,s = z1()
	var slice_one []int
	var slice_two []int
	for i := 0; i != s+1; i++ {
		if slice[i]%2==0{
			slice_one=append(slice_one, slice[i])
		} else{
			slice_two=append(slice_two,slice[i] )
		}
	}
	fmt.Println(slice_one,slice_two)
}
func main1() {
	z2()
}
