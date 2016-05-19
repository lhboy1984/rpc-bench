package main

import (
	"fmt"
	"time"
)

func SendNRequests(n int) error {
	cli, err := NewGrpcStreamClient("localhost:9090")
	if err != nil {
		return err
	}

	defer cli.Close()
	for i := 0; i < n; i++ {
		cli.Send()
	}

	return nil
}

func SendOnceByOnce(n int) {
	for i := 0; i < n; i++ {
		SendNRequests(1)
	}
}

func SendNRequestsWithMClients(m, n int) {
	fmt.Println("Start:", time.Now())
	ch := make(chan int)
	defer close(ch)

	count := 0
	for i := 0; i < m; i++ {
		go func() {
			SendNRequests(n)
			ch <- 1
		}()
	}

	for {
		_ = <-ch
		count++

		if count == m {
			break
		}
	}
	fmt.Println("End:", time.Now())
}
