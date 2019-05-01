package main

import (
    "fmt"
)

func scanNums(len int) (nums []int) {
    var num int
    for i := 0; i < len; i++ {
        fmt.Scan(&num)
        nums = append(nums, num)
    }
    return
}

func main(){
	var N int
	count := 0
	fmt.Scan(&N)
	A := scanNums(N)

	for {
		check := true
		for i,v := range A{
			if v % 2 != 0{
				check = false
			}else{
				A[i] /= 2
			}
		}
		if check{
			count++
		}else{
			break
		}
	}

	fmt.Println(count)
}