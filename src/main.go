// チャネルを用いたスレッドのラウンドロビン構成実験
//
// グローバル変数の説明
//
//	nThread: スレッドの数
//	sleepDuration: 各スレッドの処理にかかる（模擬的な）時間
package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var (
	nThread       int = 100
	sleepDuration int = 10
)

func FuncNode(label string, from chan int, to chan int) {
	for {
		v := <-from
		log.Printf("%s:%v\n", label, v)
		v += 1
		time.Sleep(time.Millisecond * time.Duration(sleepDuration))
		to <- v
	}
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	c := make([]chan int, nThread)
	for i := 0; i < nThread; i++ {
		c[i] = make(chan int, 1)
	}
	c[0] <- 0

	for i := 0; i < nThread; i++ {
		log.Printf("go FuncNode(c[%v], c[%v])\n", i, (i+1)%nThread)
		go FuncNode(fmt.Sprintf("func %v", i), c[i], c[(i+1)%nThread])
	}

	wg.Wait()
}
