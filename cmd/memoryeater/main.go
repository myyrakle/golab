package main

import (
	"fmt"
)

func main() {
	i := 0

	// GOMEMLIMIT 코드 수준 설정
	// debug.SetMemoryLimit(1000 * 1024 * 1024) // 1GB

	// GOGC 코드 수준 설정
	// debug.SetGCPercent(50) // 50%

	// 1.5gb 정도
	memorySize := 1500 * 1024 * 1024

	for {
		var bytes [][]byte

		fmt.Printf("%d번째 메모리 할당 (%d Byte)\n", i, memorySize)

		chunkSize := 1000000
		for j := 0; j < chunkSize; j++ {
			bytes = append(bytes, make([]byte, memorySize/chunkSize))
		}

		fmt.Println(bytes[len(bytes)-1][0])

		i += 1

		// 직접 할당 해제 요청
		// debug.FreeOSMemory()
	}
}
