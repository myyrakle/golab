package main

import "fmt"

func main() {
	i := 0

	// 1.3gb 정도
	memorySize := 1300 * 1024 * 1024

	var bytes [][]byte
	for {
		fmt.Printf("%d번째 메모리 할당 (%d Byte)\n", i, memorySize)

		chunkSize := 100000
		for j := 0; j < chunkSize; j++ {
			bytes = append(bytes, make([]byte, memorySize/chunkSize))
		}

		fmt.Println(bytes[len(bytes)-1][0])

		i += 1
	}
}
