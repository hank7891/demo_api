package main

import "fmt"

func main() {
	// variabletest()
	// arrayTest()
	// sliceTest()
	mapTest()
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

// 切片練習
// 陣列和切片在某些方面相似，但它們具有不同的特性和用途。在實際的程式設計中，切片更常用，因為它們提供了更多的彈性和方便的操作。
func sliceTest() {
	// 建立切片
	var slice1 []int
	fmt.Println(slice1)

	// 初始化建立長度為 5 的 int 切片
	slice2 := make([]int, 5)
	fmt.Println(slice2)

	// 初始化時也可以指定容量，切片會預留記憶體空間
	// 當擴充的元素超過容量切片會自動進行擴展
	slice3 := make([]int, 3, 5)
	slice3[0] = 1
	slice3[1] = 2
	slice3[2] = 3

	fmt.Println(slice3)

	// 切片宣告是傳址(切片指向陣列)，所以修改切片會影響到原陣列
	source := make([]int, 3)
	source[0] = 1
	source[1] = 2
	source[2] = 3

	slice4 := source[0:2]
	slice4[0] = 4
	fmt.Println(source, slice4)

	// append 可以用來擴充切片
	slice5 := []int{1, 2, 3}
	slice5 = append(slice5, 4, 5, 6)
	fmt.Println(slice5)

	// append 也可以使用其他切片來擴充，在不超過容量時，不會重新分配記憶體空間
	source = make([]int, 3, 5)
	source[0] = 1
	source[1] = 2
	source[2] = 3

	slice6 := append(source, 4, 5)
	slice6[0] = 4
	fmt.Println(source, slice6)

	// 當 append 超過容量時，會重新分配記憶體空間，並將原來的元素複製過去，此時修改切片不會影響到原陣列
	source = make([]int, 3, 5)
	source[0] = 1
	source[1] = 2
	source[2] = 3

	slice7 := append(source, 4, 5, 6)
	slice7[0] = 4
	fmt.Println(source, slice7)

	// copy 就是將一個切片的元素複製到另一個切片，故兩個切片的元素不會互相影響
	source = make([]int, 3, 5)
	source[0] = 1
	source[1] = 2
	source[2] = 3

	slice8 := make([]int, 3)
	copy(slice8, source)
	slice8[0] = 8
	fmt.Println(source, slice8)

	/**
	陣列和切片可能會有些混淆，因為它們在某些方面相似，但實際上它們有些重要的區別。

	陣列（Array）：
	 * 陣列是一個固定大小的、具有相同數據類型的元素序列。
	 * 在宣告陣列時，需要指定陣列的長度。
	 * 陣列的長度是固定的，無法改變。
	 * 陣列的宣告方式為 var arrayName [length]elementType。

	切片（Slice）：
	 * 切片是動態大小的、可以容納不同數據類型的元素序列。
	 * 切片可以根據需要動態增長或縮小，無需事先指定長度。
	 * 切片是基於陣列的一個“視圖”，它指向一個底層的陣列，但可以變長或變短。
	 * 切片的宣告方式為 var sliceName []elementType。
	 * 在你之前提供的程式碼中，使用切片字面值初始化切片的方式確實看起來類似於陣列的宣告方式，但它實際上是創建了一個切片，而不是一個固定大小的陣列。因為切片具有動態大小和彈性，所以在大多數情況下，我們更傾向於使用切片來處理集合數據，而不是使用固定大小的陣列。
	*/
}

// map 練習
func mapTest() {
	// 宣告一個 map
	var map1 map[string]int

	// 使用前需要初始化 map，不然會報錯(可以試試看把下面這行註解掉)
	// map1['a'] = 1
	map1 = make(map[string]int)
	fmt.Println(map1)

	// 當然可以宣告時同時初始化 map
	map2 := make(map[string]int)
	map2["a"] = 1
	fmt.Println(map2)

	// 宣告時同時初始化 map，並給予內容
	map3 := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(map3)

	// 刪除 map 中的元素
	delete(map3, "a")
	fmt.Println(map3)

	// 判斷 map 中是否存在某個 key
	// 這裡的 ok 是一個 bool 型態的變數，如果存在 ok 會是 true，反之則為 false
	name, ok := map3["a"]
	fmt.Println(name, ok)

	name, ok = map3["b"]
	fmt.Println(name, ok)

	/**
	map(映射) slice(切片) 陣列(陣列) 三者的區別，陣列是一個固定大小的序列，切片是基於陣列的可動態調整大小的視圖，而映射是一個鍵值對的集合。
	*/
}
