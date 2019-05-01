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
	fmt.Scan(&N)
	D := scanNums(N)
	sort.Sort(sort.Reverse(sort.IntSlice(D)))
	res := 0
	pre_mochi := 101
	for _,v := range D{
		if v < pre_mochi{
			res += 1
			pre_mochi = v
		}
	}
	fmt.Println(res)
}