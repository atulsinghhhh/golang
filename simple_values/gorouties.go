package main

import (
	"fmt"
	// "time"
	// "math/rand"
	"sync"
	// "time"
)

/*
// 	light weight threads

// 	waitgorouties

// 	Channels -> pipe in one end we sent and received
// 	communication b/w gorouties
// */

// func task(id int,w *sync.WaitGroup){
// 	defer w.Done()
// 	fmt.Println(id)

// }

// func processNum(numChan chan int){

// 	fmt.Println("processing chan",<-numChan)
// }

// func main(){

// 	// defer channel

// 	emailChan:= make(chan string,100)

// 	emailChan<-"1.com"

// 	// var wg sync.WaitGroup
// 	// for i:=0;i<10;i++{
// 	// 	wg.Add(1)
// 	// 	go task(i,&wg)
// 	// }

// 	// wg.Wait()

// 	// messageChan :=make(chan string)

// 	// messageChan <- "ping" // sent to data
// 	//  msg:=<- messageChan //received

// 	// fmt.Println(msg)

// 	// numChan:=make(chan int)
// 	// go processNum(numChan)

// 	// for {
// 	// 	numChan <- rand.Intn(100)
// 	// }

// 	// numChan <-5

// 	// time.Sleep(time.Second*2)
// }

/*
	mutex-> race condition
*/

type Post struct {
	views int
	mu sync.Mutex
}

func (p *Post) inc(wg *sync.WaitGroup){
	
	defer func(){
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock()
	p.views+=1
}

func main(){

	mypost:=Post{views: 0}
	var wg sync.WaitGroup

	for i:=0;i<100;i++{
		wg.Add(1)
		go mypost.inc(&wg)
	}

	wg.Wait()
	
	fmt.Println(mypost.views)

}

