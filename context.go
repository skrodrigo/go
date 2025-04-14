package main

import (
	"context"
	"fmt"
	"time"
)

func Context() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel() // it happens when the function returns

	go func(){
		time.Sleep(time.Second * 5)
		cancel()
	}()

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
		case <-ctx.Done():
			fmt.Println("timeout booking hotel")
		case <-time.After(time.Second * 5):
			fmt.Println("hotel booked")
		}
}