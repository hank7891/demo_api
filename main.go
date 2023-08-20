package main

import (
	"demo_api/practise"
)

func main() {
	// practise.Variabletest()
	// practise.ArrayTest()
	// practise.SliceTest()
	// practise.MapTest()
	// practise.ControlTest()
	// practise.ChannelTest()
	// practise.StructTest()
	user := &practise.User{
		Name: "hank",
		Age:  30,
	}
	user.FuncStructTest()
}
