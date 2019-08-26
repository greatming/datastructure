
def partitions(list, low, high):
    #将最左侧的值赋值给参考值k
    k = list[low]
    i = low
    j = high
    #当i下标，小于j下标的情况下，此时判断二者移动是否相交，若未相交，则一直循环
    while i < j:
        #当i对应的值小于k参考值，就一直向右移动
        while list[i] <= k:
            i += 1

        # 当j对应的值大于k参考值，就一直向左移动
        while list[j] > k:
            j = j - 1
        #若移动完，二者仍未相遇则交换下标对应的值
        if i < j:
            list[i], list[j] = list[j], list[i]

    list[low] = list[j]
    list[j] = k
    #返回k值
    return j




def quickSort(list, low, high):
    if low < high:
        k = partitions(list, low, high)
        quickSort(list, low, k -1)
        quickSort(list, k+1, high)

    return list



def main():
    data = [5,4,3,1,8,6,0,9]
    ret = quickSort(data,0, len(data)-1)
    print(ret)


if __name__ == "__main__":
    main()