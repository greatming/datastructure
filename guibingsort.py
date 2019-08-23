
def merge(left, right):
    ret = []
    i = j = 0
    while i < len(left) and j < len(right):
        if left[i] < right[j]:
            ret.append(left[i])
            i += 1
        else:
            ret.append(right[j])
            j += 1

    if i == len(left):
        ret += right[j:]
    else:
        ret += left[i:]

    return ret





def listSort(data):
    length = len(data)
    if length <= 1:
        return data

    mid = int(length / 2)
    left = listSort(data[:mid])
    right = listSort(data[mid:])
    return merge(left, right)

def main():
    data = [5,4,3,1,8,6,0,9]
    ret = listSort(data)
    print(ret)


if __name__ == "__main__":
    main()