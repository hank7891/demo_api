package main

import "fmt"

func main() {
	// variabletest()
	arrayTest()
}

// 變數練習
func variabletest() {
	// 同類型變數一起宣告
	var a, b int = 1, 2
	var c, python, java = true, false, "no!"
	fmt.Println(a, b, c, python, java)

	// 短宣告
	a1, b1 := 1, 2
	c1, python1, java1 := true, false, "no!"
	fmt.Println(a1, b1, c1, python1, java1)

	// var 混合宣告
	var (
		a2      = 1
		b2      = 2
		c2      = true
		python2 = false
		java2   = "no!"
	)
	fmt.Println(a2, b2, c2, python2, java2)
}

// 陣列練習
func arrayTest() {
	// 簡單的陣列宣告
	// 宣告一個長度為 5 的 int 陣列，預設值為 0
	var arr [5]int
	fmt.Println(arr)

	// 建造 ５個元素的 int 陣列
	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)

	// 同上，可省略陣列長度
	arr2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	// 陣列中存在不同類型的元素
	// 盡量避免使用 interface{}，會增加記憶體使用量
	arr3 := [5]interface{}{1, "three", 3, 4, 5}
	fmt.Println(arr3)

	// 查看陣列長度
	arr4 := [3]int{
		1,
		2,
		3,
	}
	fmt.Println(len(arr4))
}
