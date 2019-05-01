package main

import (
	"fmt"
	"sort"
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
	bob := 0
	alice := 0
	fmt.Scan(&N)
	A := scanNums(N)
	sort.Sort(sort.Reverse(sort.IntSlice(A)))
	for i,v := range A{
		if i % 2 == 0{
			alice += v
		}else{
			bob += v
		}
	}
	fmt.Println(alice-bob)
}