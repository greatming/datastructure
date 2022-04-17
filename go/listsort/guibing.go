package main

import "fmt"


func main ()  {
	list := []int{4,5,6,1,9,3,2}
	MergeSort(list, 0, len(list) -1)
	fmt.Println(list)
}

func MergeSort(list []int, start, end int)  {
	if start >= end {
		return
	}

	mid := (end + start) / 2
	MergeSort(list, start, mid)
	MergeSort(list, mid+1, end)
	merge(list, start, mid, end)
}

func merge(list []int, start, mid, end int)  {
	i := start 
	j := mid +1
	k := 0

	tmpArr := make([]int, end-start+1)
	for ; i <= mid && j <= end; k++ {
		if list[i] < list[j]{
			tmpArr[k] = list[i]
			i++
		}else{
			tmpArr[k] = list[j]
			j++
		}
	}
	
	for ; i <= mid; i++{
		tmpArr[k] = list[i]
		k++
	}

	for ; j <= end; j++{
		tmpArr[k] = list[j]
		k++
	}
	
	copy(list[start: end+1], tmpArr)
}