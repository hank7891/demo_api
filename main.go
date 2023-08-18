package main

import "fmt"

func main() {
	variabletest()
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
