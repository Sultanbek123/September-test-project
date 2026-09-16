package main

import "fmt"

func main() {
	fmt.Println("Hello Bitlab")
	var nums []int
	nums = append(nums, 10)
	nums = append(nums, 100)
	nums = append(nums, 1000)
	Printer(nums)
}
