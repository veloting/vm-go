package main

import (
	"flag"
	"fmt"
	"os"
)

var addr = flag.Int("addr", 0, "addr")
var pid = flag.Int("pid", 0, "pid")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	flag.Parse()
	if *addr == 0 || *pid == 0 {
		flag.Usage()
		return
	}

	fileName := fmt.Sprintf("/proc/%v/mem", *pid)
	f, err := os.OpenFile(fileName, os.O_RDWR, 0660)
	check(err)
	_, err = f.Seek(int64(*addr), 0)
	check(err)
	nop := []byte{0, 0, 0, 0}
	f.Write(nop)
	return
}
