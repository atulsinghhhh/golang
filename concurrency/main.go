package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(url string,wg *sync.WaitGroup,resultChan chan Result) {
	defer wg.Done()  // this used for cleaning function
	time.Sleep(time.Millisecond*50)

	fmt.Printf("image processed: %s\n",url)

	resultChan <- Result{
		value: url,
		err: nil,
	} // write
	// wg.Done()
}

type Result struct {
	value string
	err error
}

func main() {
	var wg sync.WaitGroup

	resultChan:=make(chan Result, 2)

	startTime:= time.Now()
	wg.Add(2)
	go worker("image_.png",&wg, resultChan)
	go worker("image2_.png",&wg,resultChan)

	wg.Wait()
	close(resultChan)

	for result := range resultChan { // read
		fmt.Printf("recived: %v\n",result);
		if result.err !=nil {
			// handle queue
		}
	}
	// fmt.Println("result",&wg)
	// fmt.Println("result",&wg)
	fmt.Printf("it took %s ms",time.Since(startTime))
	
}