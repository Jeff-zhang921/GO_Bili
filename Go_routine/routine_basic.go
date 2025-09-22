package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

//program: the code (static)
// process: when a program run, it create process:one process of program
//it will use storage, is a dynamic
//thread: program run create process, one process has muti task, each task is a thread

//one cpu work on one thread, but it can change task to do very quickly
//makes it like a concurrent, but it is cost storage and cpu

//goroutine: when execute task a you can stop in anytime to do task b
//it is a matter of logic not physic change cpu
//like in level of function using different function not changing cpu
//the fundanmental of goroutine is one thread

//inconclusion: multi threading: changing cpu fast
//     goroutine: function changing: in level of program

// ---------------------------------------------------------------------------------------------------
// main thread:hello msd
//goroutine: hello go

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello golang" + strconv.Itoa(i))

		time.Sleep(time.Second)
		//2 second output one time
	}
}

// after 130line see waitgroup
var wg sync.WaitGroup

// after line 145
var ints int

func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		ints = ints + 1
	}
}
func sub() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		ints = ints - 1
	}
}

// after line 155
// this is mutex lock:互斥锁: the actual performance is low
var sum int
var lock sync.Mutex

func add1() {
	defer wg.Done()

	for i := 0; i < 100000; i++ {
		lock.Lock() //it is important where to put the lock, to maximum speed
		sum = sum + 1
		lock.Unlock()
	}
}
func sub1() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock()
		sum = sum - 1
		lock.Unlock()
	}
}

// RWMutex:读写锁 when read goroutine is more than write goroutine,
// read don't need 互斥lock
var rwlock sync.RWMutex

func read() {
	defer wg.Done()
	rwlock.RLock()
	fmt.Println("start reading")
	//if only read, the lock won't use
	//if read and write are in same time, it will affect
	time.Sleep(time.Second)
	fmt.Println("finish reading")
	rwlock.RUnlock()
}
func write() {
	defer wg.Done()
	//the position of lock is important
	rwlock.Lock()
	fmt.Println("start writing")
	time.Sleep(time.Second)
	fmt.Println("finish writing")
	rwlock.Unlock()
}

func main() { //main thread

	go test() // open the goroutine
	for i := 0; i < 10; i++ {
		fmt.Println("hello msb" + strconv.Itoa(i))
		time.Sleep(time.Second)

	}
	//program start->main thread start execute-> go test() start go routine,(main thread still executing)
	//if main thread over, goroutine will over too no matter finish or not

	//--------------------------------------------
	//only let main thread to sleep therefore can activate the goroutine otherwise it dead directly

	//anonymous fun  c
	//go func() {
	//	fmt.Println(1)
	//}()

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second * 2)
	//when most of go routine function start to work, the loop already finish
	//so when most of goroutine capture data, i is already5, but some capture data during the loop
	//at each time the loop lanuch a goroutine,
	//it will wait until the program finish to run, but it might capture data during main thread
	//so capture data and run is in different time
	//the first goroutine that is on terminal might not be the first goroutine that generate
	//the 5th goroutine that generate will only possible outcome 4 or 5
	//there is no sequence when output goroutine
	//but you can make the number in each goroutine atleast is different

	//if sleep is first than the goroutine generate, it won't execute goroutine
	fmt.Println("----------------------------------------------------------------")
	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Println(n)
		}(i)

	}
	time.Sleep(time.Second * 2)

	fmt.Println("-----------------------------------------------------------")
	//to fix the main routine not dead before goroutine dead
	//wg is like a counter,
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done() //finish goroutine -1
			fmt.Println(i)

		}()
	}
	//it will stick until counter goes to 0,
	//you can add(total) whenyou know how many goroutine will start
	wg.Wait()
	fmt.Println("-----------------------------------------------------------")

	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(ints)
	//the result may not be 0
	//because it probabily that during the execute of sub func, add takes it, and during the execution of add func sub takes it
	//so you need a lock to make sure when a goroutine is run, other goroutine won't run
	fmt.Println("--------------------------------------------------------------------------")
	wg.Add(2)
	go add1()
	go sub1()
	wg.Wait()
	fmt.Println(sum) //this definitely print1

	fmt.Println("------------------------------------------------------------------")

	wg.Add(6)

	for i := 0; i < 5; i++ {
		go read()
	}

	go write()
	//when writing, it won't read,
	//but when no writing, it will read as if there is no lock
	wg.Wait()

}
