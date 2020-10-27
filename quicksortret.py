#参考文章:https://www.sohu.com/a/246785807_684445

def quick_sort(list, i, j):
    if  i >= j:
        return list

    low = i
    high = j
    pivot = list[i]

    while i < j:
        while i<j and list[j] >= pivot:
            j-=1
        list[i] = list[j]
        while i <j and list[i] <= pivot:
            i+=1
        list[j] = list[i]

    list[j] = pivot

    quick_sort(list, low, i-1)
    quick_sort(list, i+1, high)
    return  list













def main():
    list = [5,3,8,2,6]
    print(list)
    ret = quick_sort(list, 0, len(list)-1)
    print(ret)




if __name__ == '__main__':
    main()





