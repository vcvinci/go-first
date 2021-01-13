package _select

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

// like java future
func AsyncService() chan string {
	// todo 1 could remove ，low efficiency
	retCh := make(chan string, 1)

	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited")
	}()
	return retCh
}

func TestSelect(t *testing.T) {
	ret := <-AsyncService()
	t.Logf(ret)
}

func TestSelectTimeout(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
