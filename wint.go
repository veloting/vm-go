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

func ToMemBytes(data any) []byte {
	bf := bytes.NewBuffer([]byte{})
	if err := binary.Write(bf, binary.LittleEndian, data); err != nil {
		fmt.Println(err)
	}
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
	f.Write(ToMemBytes(int32(*pid)))
	return
}
