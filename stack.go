package main

import (
	"fmt"
	"os"
	"time"
)

func f() {
	fmt.Println("f()")
	os.Exit(0)
}

func main() {
	var a = 12345678
	for a > 0 {
		fmt.Printf("addr:%p, val:%v, pid:%v\n", &a, a, os.Getpid())
		time.Sleep(100 * time.Second)
	}

	fmt.Println("pass end")
	f()
}
