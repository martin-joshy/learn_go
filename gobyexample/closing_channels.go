package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			job, more := <-jobs
			if more {
				fmt.Println("completed job ", job)
			} else {
				fmt.Println("completed all jobs")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 4; i++ {
		jobs <- i
	}
	close(jobs)

	<-done

	_, ok := <-jobs
	fmt.Println("Received more jobs: ", ok)
}
