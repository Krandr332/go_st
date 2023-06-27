package main

import "fmt"

// Заполнение массива от от 1 до s
func z1() {
	var s int
	var slice []int
	// var number int
	fmt.Println("Введите размер массива")
	_, err := fmt.Scanln(&s)
	if err != nil {
		fmt.Println("Ошибка")
		return
	}
	fmt.Println(s)
	for i := 0; i != s+1; i++ {
		slice = append(slice, i)
		fmt.Println(slice)
	}
	return slice, s
}
func z2() {
	slice = z2()
	var slice_one []int
	var slice_two []int
	for i := 0; i != s+1; i++ {

	}
}
func main() {

}
