package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

type Profiler struct {
	memStatsBefore runtime.MemStats
	memStatsAfter  runtime.MemStats
}

func NewProfiler() Profiler {
	return Profiler{}
}

func (p *Profiler) Start() {
	// 가비지 컬렉션 실행
	runtime.GC()
	debug.FreeOSMemory()

	// 시작 전 메모리 측정
	var memStatsBefore runtime.MemStats
	runtime.ReadMemStats(&memStatsBefore)

	p.memStatsBefore = memStatsBefore
}

func (p *Profiler) Stop() (uint64, uint64) {
	// 함수 실행 후 메모리 측정
	var memStatsAfter runtime.MemStats
	runtime.ReadMemStats(&memStatsAfter)

	p.memStatsAfter = memStatsAfter

	// 사용한 힙 메모리 계산
	memoryBefore := p.memStatsBefore.HeapAlloc
	memoryAfter := p.memStatsAfter.HeapAlloc

	return memoryBefore, memoryAfter
}

func (p *Profiler) Print() {
	// 사용한 힙 메모리 계산
	memoryBefore := p.memStatsBefore.HeapAlloc
	memoryAfter := p.memStatsAfter.HeapAlloc

	// 메모리 사용량 출력
	fmt.Printf("Memory Usage: %d\n", memoryAfter-memoryBefore)
}

func main() {
	profiler := NewProfiler()

	profiler.Start()

	foo := make([]int, 500000)
	bar := make([]int, 1000000)

	foo[0] = 1
	bar[0] = 2

	profiler.Stop()

	profiler.Print()
}
