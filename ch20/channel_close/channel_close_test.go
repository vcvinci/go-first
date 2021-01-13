package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

func dataCReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			// 1. 向关闭的channel发送数据， 会导致panic
			// 2. ok true 表示正常接收， false 表示通道关闭
			// 3. 所有的channel接收者都会在channel关闭时，立即从阻塞等待中返回且上述ok值为false。
			// 这个广播机制常被利用，进行向多个订阅者同时发送信号，如退出信号
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				fmt.Println("receiver ch close! ")
				break
			}
		}
		wg.Done()
	}()
}

func TestCLoseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataCReceiver(ch, &wg)
	wg.Add(1)
	dataCReceiver(ch, &wg)
	wg.Wait()
}
