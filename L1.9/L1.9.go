package main

import (
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type channels struct {
	in  chan int
	out chan int
}

// в функции цисла считываются с канала, после чего умножаются и отправляются в канал для вывода
func multiplication(channels2 channels) {
	for i := range channels2.in {
		channels2.out <- i * 2
		time.Sleep(time.Millisecond * 100)
	}
	close(channels2.out)
}

// функция читает с канала для вывода, после чего отправляет данные на вывод
func outFd(out channels) {
	for i := range out.out {
		io.WriteString(os.Stdout, strconv.Itoa(i)+"\n")
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	chans := channels{
		// инициализируем канала, если этого не сделать, то они будут nil
		in:  make(chan int, 1),
		out: make(chan int, 1),
	}
	arr := [100]int{}
	for i := 0; i < 100; i++ {
		arr[i] = rand.Intn(600)
	}
	go multiplication(chans)
	go outFd(chans)
	for _, i := range arr {
		chans.in <- i
		time.Sleep(time.Millisecond * 100)
	}
	close(chans.in)
}
