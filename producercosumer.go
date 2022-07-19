package main

/*
模拟生产者和消费者：
生产者1：1，5，9 ...
生产者2：2，6，10 ...
生产者3：3，7，11 ...
生产者4：4，8，12 ...
消费者随机消费
*/
import (
	"sync"
)

//produce 生产者
// parameters:
// c, channel that used by the producer
// number, goroutine number, also the first number that producer produces
// wg, a wait group
// retry, max times producer produces
func produce(c chan int, number int, wg *sync.WaitGroup, retry int) {
	defer wg.Done()
	for i:=0; i< retry; i++{
		select {
		case c <- number:
			number = number + 4
		}
	}
}

//consume消费者
// parameters:
// c, channel that used by the consumer
func consume(c chan int) {
	for {
		select {
		case n := <-c:
			println(n)
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	for i:=1; i<5; i++{
		wg.Add(1)
		go produce(ch, i, wg, 3)
	}
	go consume(ch)
	wg.Wait()
}
