package main

import "fmt"

func main()  {
	list := []int{4,5,6,1,3,2}
	res := QuickSort(list, 0, len(list) - 1)
	fmt.Println(res)
}


func QuickSort(list []int, start, end int) []int {
	if start >= end{
		return list
	}
	p := partitionV2(list, start, end)
	QuickSort(list, start, p-1)
	QuickSort(list, p+1, end)
	return list
	
}

func Partition(list []int, start, end int) int{

	i := start
	pivot := list[end]
	for j :=i; j < end; j++{
		if list[j] < pivot{
			if i != j{
				list[i], list[j] = list[j], list[i]
			}
			i++
		}
	}
	list[i],list[end] = list[end], list[i] 
	return i
}


func partitionV2(arr []int, startIndex, endIndex int) int {
    var (
        pivot = arr[startIndex]
        left  = startIndex
        right = endIndex
    )

    for left != right {
        for left < right && pivot < arr[right] {
            right--
        }

        for left < right && pivot >= arr[left] {
            left++
        }

        if left < right {
            arr[left], arr[right] = arr[right], arr[left]
        }
    }

    arr[startIndex], arr[left] = arr[left], arr[startIndex]

    return left
}
