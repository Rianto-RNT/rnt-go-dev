//--Summary:
//  Create a program that utilizes goroutines to run the provided calculation
//  function on a number of jobs. The results from the goroutines must be
//  communicated back to the main thread using a channel, and then added
//  together.
//
//--Requirements:
//* Run `longCalculation` for each job generated by the `makeJobs` function
//* Each job must be run in a separate goroutine
//* The result from `longCalculation` must be provided to the main function
//  using a channel
//* Sum the results from each job to generate a final result, and print it
//  to the terminal

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Job int

func longCalculation(i Job) int {
	duration := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(duration)
	fmt.Printf("Job %d complete in %v\n", i, duration)
	return int(i) * 30
}

func makeJobs() []Job {
	jobs := make([]Job, 0, 100)
	for i := 0; i < 100; i++ {
		jobs = append(jobs, Job(rand.Intn(1000)))

	}
	return jobs
}

func runJob(resultChan chan int, i Job) {
	resultChan <- longCalculation(i)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	jobs := makeJobs()
	resultChan := make(chan int, 10)

	for i := 0; i < len(jobs); i++ {
		go runJob(resultChan, jobs[i])
	}

	resultCount := 0
	sum := 0
	for {
		result := <-resultChan

		sum += result
		resultCount += 1
		if resultCount == len(jobs) {
			break
		}
	}
	fmt.Println("Sum is", sum)
}
