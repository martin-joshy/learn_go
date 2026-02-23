package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Println("Started working on job - ", job, "worker id: ", id)
		time.Sleep(time.Second)
		fmt.Println("Finished working on job - ", job, " worker id: ", id)
		result <- 2 * job
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	result := make(chan int, numJobs)

	for i := 0; i <= 3; i++ {
		go worker(i, jobs, result)
	}

	for j := 0; j <= 5; j++ {
		jobs <- j
	}

	close(jobs)

	for r := 0; r <= numJobs; r++ {
		<-result
	}
	fmt.Println("All work has been completed")
}
