package main

import (
    "fmt"
)

func main(){
	var A,B,C,X int
	count := 0
	fmt.Scan(&A,&B,&C,&X)
	for a:=0;a<=A;a++{
		for b:=0;b<=B;b++{
			for c:=0;c<=C;c++{
				if a*500 + b*100 + c*50 == X{
					count++
				}
			}
		}
	}
	fmt.Println(count)
}