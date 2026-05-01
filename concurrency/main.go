package main

import (
	"fmt"
	"time"
)

func worker(url string) string{
	time.Sleep(time.Millisecond*50)

	fmt.Printf("image processed: %s\n",url)
	return url
}

func main() {
	startTime:= time.Now()
	result1:=worker("image_.png")
	result2:=worker("image2_.png")

	fmt.Println("result",result1)
	fmt.Println("result",result2)
	fmt.Printf("it took %s ms",time.Since(startTime))
	
}