package main

import (
"log"
"time"
)

func dummy(c chan int) {
time.Sleep(3 * time.Second)
random := <-c
log.Println("receiving--", random)
}

func main(){
ch := make(chan int)
go dummy(ch)
log.Println("sending--")
ch <- 100
}