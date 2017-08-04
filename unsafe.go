package main

import "fmt"
import "unsafe"

type packet struct {
	opcode uint16
	data   [1022]byte
}

type file_info struct {
	file_size uint32     // 4 bytes
	file_name [1018]byte //this struct has to fit in packet.data
}

func makeData() []byte {
	fi := file_info{file_size: 1 << 20}
	copy(fi.file_name[:], []byte("test.x64"))
	p := packet{
		opcode: 1,
		data:   *(*[1022]byte)(unsafe.Pointer(&fi)),
	}
	mem := *(*[1022]byte)(unsafe.Pointer(&p))
	return mem[:]
}
func main() {
	data := makeData()
	fmt.Println(data)
	p := (*packet)(unsafe.Pointer(&data[0]))
	if p.opcode == 1 {
		fi := (*file_info)(unsafe.Pointer(&p.data[0]))
		fmt.Println(fi.file_size, string(fi.file_name[:8]))
	}
}

