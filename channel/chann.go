package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// but if you give space to channel (data := make chan int 2) , so that time you can insert (its called bufferd channel )\

//channels are mainly used multi thread concepts



func channelBaicTest() {

	// when we insert a data into the channel at the same time it need to exit , otherwise it show dead lock 

	dataChan := make(chan int )

	go func() {
      dataChan <- 20
	}()
	n := <-dataChan

	fmt.Println(n)

}


func doWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)

}

func main() {


	// with out close(dataChann)the channel we can see the deadlock , channcels are waiting for the data 

  // close(dataChan)-> when i close channel its fine 
	dataChan := make(chan int )




	go func() {


		wg := sync.WaitGroup{}
     for i:= 0; i< 100;i++{

       wg.Add(1)

         go func() {
            defer wg.Done()
		result := doWork()

		dataChan <- result 
		 
	} ()

	 }

	 wg.Wait()
	 close(dataChan)
	}()

	for n := range dataChan{

		

		fmt.Println("dataChan->",n)
		
	}

	

}