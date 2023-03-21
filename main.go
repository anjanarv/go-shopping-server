package main

import (
"log"
)

func dummy(c chan int) {
random := <-c
log.Println("receiving--", random)
}

func main(){
ch := make(chan int)
go dummy(ch);
log.Println("sending--")
ch <- 100;
}