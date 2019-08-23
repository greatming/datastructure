






def insertSort(arr):
    length = len(arr)
    for i in range(1,length):
        x = arr[i]
        for j in range(i,-1,-1):
            # j为当前位置，试探j-1位置
            if x < arr[j-1]:
                arr[j] = arr[j-1]
            else:
                # 位置确定为j
                break
        arr[j] = x


def main():
    data = [5,4,3,1,8,6,0,9]
    insertSort(data)
    print(data)


if __name__ == "__main__":
    main()