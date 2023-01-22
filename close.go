package main

import "fmt"

var int_numbers chan int
var done chan bool


func receive(){
	for{
		num, next := <- int_numbers
		if next {
			fmt.Println("received: ", num)
		} else {
			fmt.Println("all numbers received. Thank you :-)")
			done <- true
			return
		}
	}
}

func main() {
	int_numbers = make(chan int, 35)
	done = make(chan bool)

	go receive()

	for i := 0 ; i <= 30 ; i++ {
		int_numbers <- i
	}
	close(int_numbers)
	
	<- done

}