package main

import (
	"fmt"
	"os"
)

func main(){
	var N,Y int
	fmt.Scan(&N,&Y)
	for i:=0;i<=N;i++{
		for j:=0;j<=N;j++{
			if i + j > N{
				break
			}
			money := 10000 * i + 5000 * j + 1000 * (N - i - j)
			if money == Y{
				fmt.Println(i,j,N-i-j)
				os.Exit(0)
			}
		}
	}
	fmt.Println(-1,-1,-1)
}