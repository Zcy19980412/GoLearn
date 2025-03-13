package main

import (
	"fmt"
	"time"
)

func main() {

	//并行：多核cpu同时执行多个指令
	//并发：cpu快速调度，形成并行的假象

	//同步：任务按顺序执行
	//异步：任务可以启动后继续执行其他任务

	//多个控制流，共同操作同一个共享资源的情况，都需要同步
	//锁：锁是加在共享资源上的，只有拿到锁的才能对共享资源进行操作
	//互斥锁：建议锁（线程可以不拿锁直接读写，需要人为的区拿锁），人为控制锁住期间的线程行为，同一时刻只能被一个线程持有
	//读写锁：一把锁具有读属性和写属性，读锁可以被多个线程同时持有，写锁统一时间只能被一个线程持有。对于互斥锁进行了优化。

	//为什么设计协程：
	/**
	创建开销	高（MB 级栈）	低（2KB 初始栈）
	调度方式	内核态（OS 负责）	用户态（Go 运行时调度）
	上下文切换	高（保存寄存器、堆栈等）	低（轻量级调度）
	最大并发数	受 OS 线程数限制（几千）	数十万甚至百万
	同步机制	需要锁（Mutex）	推荐使用 Channel
	阻塞处理	线程阻塞会占用 CPU	Goroutine 阻塞时自动调度
	*/
	testGoroutine()

}

func testGoroutine() {
	go sing()
	go dance()

	//主go程结束，子go程终止
	for {

	}
}

func sing() {
	for i := 0; i < 10; i++ {
		fmt.Println("sing")
		time.Sleep(1 * time.Second)
	}
}

func dance() {
	for i := 0; i < 10; i++ {
		fmt.Println("dance")
		time.Sleep(1 * time.Second)
	}
}
