package cmd

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/exoscale/egoscale"
	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

type task struct {
	egoscale.AsyncCommand
	string
}

type taskStatus struct {
	id        int
	jobStatus egoscale.JobStatusType
}

// asyncTasks message variable must have same size with cmds
func asyncTasks(tasks []task) ([]interface{}, []error) {

	//init results
	e := make([]error, 0, len(tasks))
	errors := &e
	r := make([]interface{}, len(tasks))
	responses := &r

	//create task Progress
	taskBars := make([]*mpb.Bar, len(tasks))
	maximum := 10
	var wg sync.WaitGroup
	p := mpb.New(mpb.WithWaitGroup(&wg), mpb.WithContext(gContext), mpb.WithWidth(40))
	wg.Add(len(tasks))

	//exec task and init bars
	for i, task := range tasks {
		c := make(chan taskStatus)
		go execTask(task, i, c, responses, errors)
		taskBars[i] = p.AddBar(int64(maximum),
			mpb.PrependDecorators(
				// simple name decorator
				decor.Name(task.string),
				// decor.DSyncWidth bit enables column width synchronization
				decor.Percentage(decor.WCSyncSpace),
			),
			mpb.AppendDecorators(
				// replace ETA decorator with "done" message, OnComplete event
				decor.OnComplete(
					// ETA decorator with ewma age of 60
					decor.EwmaETA(decor.ET_STYLE_GO, 60), "done",
				),
			),
		)

		//listen for bar progress
		go func(chanel chan taskStatus) {
			defer wg.Done()
			defer close(chanel)
			max := 100 * time.Millisecond
			count := 1
			for status := range chanel {
				start := time.Now()
				time.Sleep(time.Duration(rand.Intn(10)+1) * max / 10)
				if status.jobStatus == egoscale.Pending {
					if count < maximum {
						taskBars[status.id].IncrBy(1, time.Since(start))
					}
				} else {
					taskBars[status.id].IncrBy(maximum, time.Since(start))
					return
				}
				count++
			}
		}(c)
	}

	p.Wait()

	if len(*errors) > 0 {
		return nil, *errors
	}
	return *responses, nil
}

func execTask(task task, id int, c chan taskStatus, resps *[]interface{}, errors *[]error) {
	response := cs.Response(task.AsyncCommand)
	var errorReq error
	cs.AsyncRequestWithContext(gContext, task.AsyncCommand, func(jobResult *egoscale.AsyncJobResult, err error) bool {
		if err != nil {
			errorReq = err
			return false
		}

		if jobResult.JobStatus == egoscale.Success {
			if errR := jobResult.Result(response); errR != nil {
				errorReq = errR
				return false
			}
			(*resps)[id] = response
			c <- taskStatus{id, egoscale.Success}
			return false
		}

		c <- taskStatus{id, egoscale.Pending}
		return true
	})

	if errorReq != nil {
		c <- taskStatus{id, egoscale.Failure}
		*errors = append(*errors, errorReq)
	}
}

// asyncRequest if no response expected send nil
func asyncRequest(cmd egoscale.AsyncCommand, msg string) (interface{}, error) {
	response := cs.Response(cmd)

	fmt.Print(msg)
	var errorReq error
	cs.AsyncRequestWithContext(gContext, cmd, func(jobResult *egoscale.AsyncJobResult, err error) bool {

		fmt.Print(".")

		if err != nil {
			errorReq = err
			return false
		}

		if jobResult.JobStatus == egoscale.Success {
			if errR := jobResult.Result(response); errR != nil {
				errorReq = errR
				return false
			}

			fmt.Println(" success.")
			return false
		}

		return true
	})

	if errorReq != nil {
		fmt.Println(" failure!")
	}

	return response, errorReq
}
