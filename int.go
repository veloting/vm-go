package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var a = 12345678
	for a > 0 {
		fmt.Printf("addr:%p, val:%v, pid:%v\n", &a, a, os.Getpid())
		time.Sleep(1 * time.Second)
	}

	fmt.Println("pass end")
}
