package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func BackupWorker(bkpContext context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	etcSec := 4
	for i := 1; i <= etcSec; i++ {
		select {
		case <-bkpContext.Done():
			fmt.Println("Backup process interrupted due to context, reason:", bkpContext.Err())
			return
		default:
			fmt.Printf("Backup Process Done %d%%", (100/etcSec)*i)
			fmt.Println()
		}
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Backup Process Completed Successfully")
}

func main() {
	var wg sync.WaitGroup
	bkpContext, bkpCancleFn := context.WithTimeout(context.Background(), 3*time.Second)
	defer bkpCancleFn()

	wg.Add(1)
	go BackupWorker(bkpContext, &wg)

	wg.Wait()
	fmt.Println("main: Exit")
}
