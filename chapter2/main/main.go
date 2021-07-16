package main

import "fmt"


func main() {
	var example = [5]int32{6, 7, 843, 43, 23}
	insertionSort(&example)
	fmt.Println(example)
}


