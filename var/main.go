package main

import (
	"errors"
	"fmt"
)

type testStruct struct {
	a int
}

func getError2() (num int, err error){
	return 3, errors.New("Try this")
}

func getError() (err error){
	if a, err := getError2(); err != nil{
		fmt.Println(a)
	}

	return
}

func main(){
	err:=getError()
	if err != nil{
		fmt.Println(err.Error())
	}
	var test1 testStruct
	fmt.Println(test1)
	var test2 *testStruct
	fmt.Println(test2)
}
