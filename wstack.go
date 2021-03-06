package main

import (
	"bytes"
	"encoding/binary"
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

func ToMemBytes(data []byte) []byte {
	bf := bytes.NewBuffer([]byte{})
	binary.Write(bf, binary.LittleEndian, data)
	return bf.Bytes()
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
	nop := ToMemBytes([]byte{0x80, 0x89})
	f.Write(nop)
	return
}
