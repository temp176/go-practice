package main

import (
    "fmt"
)

func digit_sum(n int)int{
	res := 0
	dig := 0
	for n > 0{
		dig = n%10
		res += dig
		n /= 10
		if n == 0{
			break
		}
	}
	return res
}

func main(){
	var N,A,B int
	res := 0
	fmt.Scan(&N,&A,&B)
	for n:=0;n<=N;n++{
		if v := digit_sum(n); v >= A && v <= B{
			res += n
		}
	}
	fmt.Println(res)
}