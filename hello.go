package main

import "fmt"

func main() {
	fmt.Println("Aloha")

	var age = 2
	var num = 2.35648
	var str = "Hello world"
		fmt.Println (age, num, str)

	var a = 10
	var b = 13
	var res0 int
	var res1 int
	var res2 int
	var res3 int
	//var res4 int
	var res5 int
	res0 = a + b
	res1 = a-b
	res2 = a * b
	res3 = b / a
	// res4 = b % a
	res5 = 456789
	if age < 5 {
		fmt.Println("Сопляк!")
	}
	fmt.Println("result is", res0)
	fmt.Println("result is", res1)
	fmt.Println("result is", res2)
	fmt.Println("result b/a is", res3,"Где а = ", a,",а b =", b)
	fmt.Println("result is", res5)
}
