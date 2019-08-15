package main

import "fmt"

type C1Info struct {
	Id int
}

func testSliceNil(){
	var C1Infos []*C1Info
	fmt.Println(C1Infos)
	C1Infos = nil
	fmt.Println(C1Infos)
	C1Infos = make([]*C1Info, 0)
	fmt.Println(C1Infos)
	for _, v :=range C1Infos{
		fmt.Println("hello world", v)
	}
	fmt.Println(C1Infos)
}

func main(){
	testSliceNil()
}
