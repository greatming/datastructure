package main

import "fmt"


func InsertSort(list []int) []int{
	cnt := len(list)
	if cnt <= 1{
		return list
	}
	for i := 1; i < cnt; i++{
		val := list[i]
		j := i-1
		for; j >=0; j--{
			if val >= list[j]{
				break
			}
			list[j+1] = list[j]
		}
		list[j+1] = val
	}
	return list
}


func main()  {
	list := []int{4,5,6,1,3,2}
	res := InsertSort(list)
	fmt.Println(res)
}