package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type readChannel chan interface{}

// reader - функция считывания данных из канала
func reader(ctx context.Context, channel readChannel) {
	for {
		select {
		// Как только канал закрыт, то выходим
		case <-ctx.Done():
			return
			// Как только данные пришли, то считываем их
		case data := <-channel:
			fmt.Println("Func read", data, "in", time.Now().UTC())
		}
	}
}

// send - функция, которая отправляет данные в канал
func send(ctx context.Context, channel readChannel) {
	for {
		select {
		// Как только канал закрыт, то выходим
		case <-ctx.Done():
			return
			// В обычном состоянии записываем данные в канал
		default:
			time.Sleep(time.Millisecond * 100)
			channel <- rand.Intn(1000)
		}
	}
}

func main() {
	readChan := make(readChannel, 1)
	var numberSec time.Duration = 3
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*numberSec)
	defer cancel()
	go reader(ctx, readChan)
	send(ctx, readChan)
	fmt.Println("All data has been sent!")
}
