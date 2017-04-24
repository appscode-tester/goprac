package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	for i:=10; i>0; i-- {
		fmt.Println("Current time: ", t)
		time.Sleep(time.Second*20)
	}
}
