package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	sum      int64
	upperlim int
)

func nqueen(n int) string {
	now := time.Now()

	uint8N := uint8(n)
	sum = 0
	upperlim = 1<<uint8N - 1
	gotest(0, 0, 0, uint8N)

	return fmt.Sprintf("%d queue:\t%d\n时间:\t%v", n, sum, time.Since(now))
}

func gotest(row, ld, rd int, n uint8) {
	pos := upperlim & ^(row | ld | rd)
	wg := new(sync.WaitGroup)
	for i := 0; i < int(n/2); i++ {
		p := pos & (-pos)
		pos -= p
		wg.Add(1)
		go func() {
			test(row+p, (ld+p)<<1, (rd+p)>>1)
			wg.Done()
		}()
	}
	// 如果是奇数，还需要在算一层
	if n%2 != 0 {
		p := pos & (-pos)
		pos -= p
		wg.Add(1)
		go func() {
			testOdd(row+p, (ld+p)<<1, (rd+p)>>1)
			wg.Done()
		}()
	}
	wg.Wait()
}

func test(row, ld, rd int) {
	if row != upperlim {
		pos := upperlim & ^(row | ld | rd)
		for pos != 0 {
			p := pos & (-pos)
			pos -= p
			test(row+p, (ld+p)<<1, (rd+p)>>1)
		}
	} else {
		atomic.AddInt64(&sum, 2)
	}
}

func testOdd(row, ld, rd int) {
	if row != upperlim {
		pos := upperlim & ^(row | ld | rd)
		for pos != 0 {
			p := pos & (-pos)
			pos -= p
			testOdd(row+p, (ld+p)<<1, (rd+p)>>1)
		}
	} else {
		atomic.AddInt64(&sum, 1)
	}
}
