package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

// 当心recover 称为恶魔
// 形成僵尸进程 导致health check 失效
// "Let it Crash!"
func TestPanicVxExit(t *testing.T) {
	/*defer func() {
		fmt.Println("Finally!!")
	}()*/
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()
	fmt.Println("Start")
	//os.Exit(-1)
	panic(errors.New("Something wrong!"))
}
