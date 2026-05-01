package main

import (
	"fmt"
	"sync"
	"time"
)



func worker(jobsChan chan string,wg *sync.WaitGroup,resultChan chan Result) {
	defer wg.Done()  // this used for cleaning function

	for result := range jobsChan { // read
		time.Sleep(time.Millisecond*50)
		fmt.Printf("image processed: %s\n",result)
		// fmt.Printf("recived: %v\n",result);

		resultChan <- Result{
			value: result,
			err: nil,
		} //
		
	}
	

	fmt.Println("Worker shutting down");

	// resultChan <- Result{
	// 	value: url,
	// 	err: nil,
	// } // write
	// wg.Done()
}

type Result struct {
	value string
	err error
}

func main() {

	jobs := []string{
		"image_1.png",
		"image_1.png",
		"image_1.png",
		"image_1.png",
		"image_1.png",
		"image_1.png",
		"image_1.png",
		"image_1.png",
		"image_1.png",
		"image_1.png",
		"image_1.png",
		"image_1.png",
		"image_1.png",
		"image_1.png",
		"image_1.png",
		"image_1.png",
		"image_1.png",
		"image_1.png",
	}
	var wg sync.WaitGroup

	// total not workers
	totalWorkers:= 7

	resultChan:=make(chan Result, 50)
	jobsChan:= make(chan string,len(jobs))

	// keeps listening for jobs until the channel is closed
	for i:=1;i<=totalWorkers;i++{
		wg.Add(1)
		go worker(jobsChan,&wg,resultChan)
	}


	startTime:= time.Now()
	// wg.Add(2)

	// for _,job :=range jobs {
	// }
	// go worker("image_.png",&wg, resultChan)
	// go worker("image2_.png",&wg,resultChan)

	go func () {
		wg.Wait()
		close(resultChan)

	} ()

	// send the jobs
	for i:=0;i<len(jobs);i++{
		jobsChan <-jobs[i]
	}

	close(jobsChan)

	for result := range resultChan { // read
		fmt.Printf("jobs completed: %v\n",result);
		if result.err !=nil {
			// handle queue
		}
	}
	// fmt.Println("result",&wg)
	// fmt.Println("result",&wg)
	fmt.Printf("it took %s ms",time.Since(startTime))
	
}