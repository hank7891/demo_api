package practise

import (
	"fmt"
	"time"
)

// 變數練習
func Variabletest() {
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
func ArrayTest() {
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
func SliceTest() {
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
func MapTest() {
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

// 控制結構練習
func ControlTest() {
	// if
	a := 1
	if a == 1 {
		fmt.Println("a == 1")
	} else {
		fmt.Println("a != 1")
	}

	// if 也可以在判斷前先執行一個語句
	if b := 1; b == 1 {
		fmt.Println("b == 1")
	} else {
		fmt.Println("b != 1")
	}

	// else if
	if c := 2; c == 1 {
		fmt.Println("c == 1")
	} else if c == 2 {
		fmt.Println("c == 2")
	} else {
		fmt.Println("c != 1 && c != 2")
	}

	// if 可以省略判斷式，當判斷式為 true 時執行
	if true {
		fmt.Println("Is true")
	}

	// switch
	switch i := 1; i {
	case 1:
		fmt.Println("1")
	case 2, 3:
		fmt.Println("2 or 3")
	case 4:
		fallthrough
	case 5:
		fmt.Println("5")
	}

	// switch 比對
	i := 5
	switch {
	case i > 5:
		fmt.Println(i, "> 5")
	case i < 5:
		fmt.Println(i, "< 5")
	case i == 5:
		fmt.Println(i, "= 5")
	}

	// for
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for 可省略前後語句
	j := 1
	for j < 5 {
		j += j
	}
	fmt.Println("for 測試")
	fmt.Println(j)
}

// 通道練習
// 通道（Channel）是一種用於在不同的 Go 協程（Goroutine）之間進行通訊和同步的機制。通道提供了一種安全的方法來傳遞數據，確保協程之間的同步和互斥。
func ChannelTest() {
	// 建立一個通道
	ch := make(chan int)

	// 通道的寫入和讀取都是阻塞的，所以必須在不同的協程中進行
	go func() { ch <- 1 }()

	// 從通道中讀取數據
	fmt.Println(<-ch)

	// channel 有個 select 應用，可以讓 channel 有非同步的效果
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- 2
	}()

	select {
	case num := <-ch1:
		fmt.Println("Received from ch1:", num)
	case num := <-ch2:
		fmt.Println("Received from ch2:", num)
	}

	// 一個 select 僅會監聽一次，可以使用 for 迴圈讓 select 監聽多次
	// 記得註解上方單次監聽的 select
	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case num := <-ch1:
	// 		fmt.Println("Received from ch1:", num)
	// 	case num := <-ch2:
	// 		fmt.Println("Received from ch2:", num)
	// 	}
	// }
}

// 結構練習
func StructTest() {
	type User struct {
		Name string
		Age  int
	}

	// 宣告一個 User
	user := User{
		Name: "hank",
		Age:  30,
	}

	// 取得 User 的屬性
	fmt.Println(user.Name, user.Age)
}

// 函示串接結構練習
type User struct {
	Name string
	Age  int
}

func (user *User) FuncStructTest() {
	fmt.Println(user.Name, user.Age)
}

// defer 練習
// defer 會在函式結束前執行，可以用來釋放資源
func DeferTest() {
	// 可以看到雖然 defer 在最前面，但是會在函式結束前執行，所以會先印出 hello
	defer fmt.Println("world")
	fmt.Println("hello")

	/**
	因為以上特性 defer 通常用在。
	 * 資源釋放
	 * 記錄紀錄
	 * 解鎖操作
	 * 執行結束後操作
	 * 清理資源
	*/
}

// defer 製造 try catch 練習
func DeferTryCatchTest() {
	// defer 會在函式結束前執行，所以可以用來製造 try catch 的效果
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("Start")
	panic("Oops! An error occurred")
	fmt.Println("hello")

	/**
	特別介紹 panic 和 recover 兩個內建函示。
	 * panic 會中斷程式，並且執行 defer，所以可以用來製造 try catch 的效果。
	 * recover 可以捕捉 panic 的錯誤，並且讓程式繼續執行。
	*/
}
