package main

import (
	"fmt"
	"sync"
	"time"
)

//even if you add lock, it still don't guarantee the sequence

//channel property
//channel is a datatype

// channel is a FIFO type(Queue)
// if many go routine operate on same channel, they don't need to add lock
//attention! goroutine operate on channel, not inside channel
//only one type can in one channel
// var name chan datatype

var lock sync.RWMutex
var wg sync.WaitGroup

func main() {
	//define the channel
	var intChan chan int
	//you can use make to initialise
	//make is only use for slice map and channel
	intChan = make(chan int, 3) //you can put 3 datatype of int
	fmt.Printf("intChan value:%v\n", intChan)
	//channal is the address

	//put data in to channel
	intChan <- 10
	num := 20
	intChan <- num
	fmt.Printf("the actual length is:%v, the capacity is%v\n", len(intChan), cap(intChan))

	//cannot put more than the capacity
	//read data from channel
	num1 := <-intChan
	fmt.Println("num1 val:", num1)

	//to close the channel
	//after close, it can read but cannot write
	close(intChan)
	num2 := <-intChan
	fmt.Println("num2 val:", num2)

	//loop through the channel
	var chanInt chan int
	chanInt = make(chan int, 100)
	for i := 0; i < 100; i++ {
		chanInt <- i
	}
	//getting the data from channal can only use for range
	//there is no index inside channel so only one
	//if you don't close the channel before loop through, deadlock!
	//go don't know when to stop because it think that more stuff may comes in

	close(chanInt)
	for val := range chanInt {
		fmt.Println("Value:", val)
	}

	//-------------------------------------------------
	//goroutine work with channel
	intChan2 := make(chan int, 50)
	wg.Add(2)
	go writedata(intChan2)

	go readdata(intChan2)
	wg.Wait()
	//---------------------------------------------------------------------------------------------------------------------
	//in default, the channel is read write
	//it but it can only read or write too
	var intchan chan<- int //only write
	intchan = make(chan int, 5)
	intchan <- 10
	fmt.Println("intchan:", intchan)
	// num:=<-intchan
	var intchan2 <-chan int //only read
	if intchan2 != nil {
		num2 := <-intchan2
		fmt.Println(num2)

	}

	//blockage in channel
	//when channel only write no read, it will error(no read no new room)
	//---------------------------------------------------------------------------------------------------------------------------
	//select: fix the select problem if have many channel at once: choose one in fair
	//define a int ,define a string
	fmt.Println("----------------------------------------------------------------------------------------------")
	intChan3 := make(chan int, 1)
	stringChan := make(chan string, 1)
	go func() {
		time.Sleep(time.Second)
		intChan3 <- 10
	}()
	go func() {
		time.Sleep(time.Second * 2)
		stringChan <- "Hello"
	}()
	//when getting data it need to take
	//fmt.Println(<-intChan3)
	//fmt.Println(<-stringChan)

	//choose one to execute who first get the data execute who
	select { //after case need to be an IO operate
	case v := <-intChan3:
		fmt.Println("intchan:", v)
	case v := <-stringChan:

		fmt.Println("stringchan:", v)
	//With default, it becomes non-blocking (if no data ready, run default).
	default:
		fmt.Println("default")
	}
	//--------------------------------------------------------------------------------------------------------------------------------------------
	//when go routine hit sleep func, that really under hood
	//The scheduler marks it as “waiting” (until 2s pass).
	//While it’s waiting, the scheduler gives CPU to other goroutines.
	//After 2s, the sleeping goroutine becomes runnable again, and the scheduler may resume it.

	//main don't have the previlage to run first unless it is the start of the program(which only have main to run)

	//------------------------------------------------------------------------------------------------------------------------------------------------------
	go printnum()
	go devide()
	time.Sleep(time.Second)
}

//reader will only not reader when find the channel close.
//if you close in loop, it will stop read immediately
//the reason why deadlock is occur is that there is no close and read will never stop

//channel property
//channel don't decide which function run first,but itself has lock function
//reader might execute before writer, but reader need to wait until writer send data,
//the last part of sentence is the property of channel

//Unbuffered channel forces strict synchronization (send ↔ receive).
//so it always start with send, when no send, receive will be block, when no receive , next send will be block
//this is the property of channel

//Buffered channel allows senders/readers to run ahead, but only if the producer/consumer don’t pause.

func writedata(intChan2 chan int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		//	lock.Lock()
		intChan2 <- i

		fmt.Println("the data have been write is:", i)
		time.Sleep(time.Second)

	}
	close(intChan2)

}
func readdata(intChan2 chan int) {
	defer wg.Done()
	for val := range intChan2 {

		fmt.Println("the data have been read", val)
		time.Sleep(time.Second)

	}
}

// ------------------------------------------------------------------------
func printnum() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// this is actually a wrong goroutine
// use a defer recover to a goroutine that might goes wrong
func devide() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("devide error:", err)
		}
	}()
	num1 := 10
	num2 := 0
	fmt.Println(num1 / num2)
}
