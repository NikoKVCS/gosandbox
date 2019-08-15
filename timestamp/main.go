package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix())
	time.Sleep(time.Second * 2)
	fmt.Println(time.Now().Unix())
	time.Sleep(time.Second * 2)
	fmt.Println(time.Now().Unix())
}
