package groutine__test

import (
	"fmt"
	"testing"
	"time"
)

/**
 * default stack size:
 * java 1M
 * go   2K
 * KSE (kernel Space Entity) 的对应关系
 * java Thread  1:1
 * Groutine     M:N
 *
 * 理解groutine:
 * 线程切换 射击到内核对象的切换 是一个很大的消耗
 *
 * M : System Thread -> KSE
 * P : Processor
 * G : Goroutine
 *
 * 没个P都有一个G队列，依次运行。 如果一个G调用时间过长，是不是会延迟很久？
 * go起来后会有一个守护线程，去计数没个P运行完成G的数量，一段时间，发现某个P的数量没有变的时候，
 * 会向G的任务栈里放一个特别标记，当这个G运行遇到非内联函数会读到这个标记，就会中断插到等候G队列的队伍，然后切换别的G继续运行
 *
 * 另一个提高并发的机制：
 * 当某一个G被系统中断了，可能是IO，需要等待，为了提高整体的并发
 * 会把P移到另一个可使用的系统线程当中继续执行P所挂的队列里的其他的G
 * 当中断的G被唤醒，完成之后，会把它加到P的某一个的等待队列里或是全局等待队列当中
 * 需要注意的是：
 * 当一个G被中断的时候，在寄存器里的运行状态也会保存到这个G对象里，
 * 当G在此获得运行机会的时候，会重新写入寄存器重新运行
 */
func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)

		// i 被共享了
		/*go func() {
			fmt.Println(i)
		}()*/
	}
	time.Sleep(time.Microsecond * 50)
}
