package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/**
	Go 只保留最核心的 互斥锁 和 读写锁。

	首先涉及到共享数据的并发修改，那么一定要保证原子性，通过atomic或者加锁（mutex）的方式。
	如果在共享数据的并发修改基础上还涉及共享数据的并发访问，那么需要用到读写锁。

	Go 的设计哲学是：

	用 Channel 解决大部分并发问题，尽量减少锁的使用。
	锁的设计要足够简单，以便避免 Java 复杂的锁管理问题（死锁、活锁、饥饿等）。
	不要引入 Java 那种灵活但复杂的锁机制，而是提供最基础的工具，让开发者自己组合使用。

	共享数据的并发修改 → 必须保证原子性 (sync.Mutex or sync/atomic)
	共享数据的并发访问 → sync.RWMutex 允许读并发，写互斥
	如果是简单数值变量的并发修改，可以用 sync/atomic


	什么时候用 sync.Mutex，什么时候用 sync.RWMutex？
	如果读操作远多于写操作（典型的查询业务，如银行账户查询）
	→ 使用 sync.RWMutex，让读操作可以并发执行，提升性能。
	如果读写频率差不多，或者写操作很多（比如交易处理系统）
	→ 直接用 sync.Mutex，避免 RWMutex 频繁升级为写锁的开销。
	*/

	//Channel 死锁  Mutex 死锁
	//mutex

	//testDeadLock()
	//testMutex()
	testRWMutex()
}

func testDeadLock() {
	//死锁是一种状态，指多个协程或进程因互相等待对方释放资源而永远无法继续执行的情况。
	// all goroutines are asleep - deadlock!

	c := make(chan int)
	c <- 2
	fmt.Println(<-c)
}

func Printer(s string) {
	for i := range s {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%c", s[i])
	}
}

func testMutex() {
	//粗颗粒度的锁，不分读写
	//Mutex ： 互斥锁，互斥量。  是建议锁！！  建议锁指这个锁不强制加在某个方法或者代码块上，而是有开发者自己管理，可以不用
	var printerMutex sync.Mutex = sync.Mutex{}

	go func() {
		printerMutex.Lock()
		Printer("worker")
		printerMutex.Unlock()
	}()
	go func() {
		printerMutex.Lock()
		Printer("leader")
		printerMutex.Unlock()
	}()

	for {
	}

}

func testRWMutex() {
	//细颗粒度的锁：覆盖率互斥锁的功能，添加了读锁，特点是在读锁被释放时，写锁可以并发，否则阻塞
	//假如有一个学校在招生，多个窗口同时录入学生，每招一个学生，统计数据： 总学生人数，全校总人数。都应该被+1。 这一篮子操作应具有原子性。
	//如果仅使用 写锁锁住 {总学生人数+1; 全校总人数+;}可以保证最后的人数总是正确的。  但是过程中的读取，比如总学生人数+1后 两列数据被读取
	//则出现错误。 所以  读取不应该在写入过程中。 故而产生了读锁。

	//10个窗口同时录入学生，1个窗口录入10个

	var count1Int = 0
	var count2Int = 0
	count1 := &count1Int
	count2 := &count2Int

	for i := 0; i < 10; i++ {
		//go func() {
		//	for range 10 {
		//		count1++
		//		time.Sleep(100 * time.Millisecond)
		//		count2++
		//	}
		//}()
		go increase(count1, count2)
	}

	<-time.After(5 * time.Second)
	fmt.Println(*count1, *count2)

	for {
	}

}

func increase(count1, count2 *int) {
	for range 10 {
		*count1++
		time.Sleep(100 * time.Millisecond)
		*count2++
	}
}
