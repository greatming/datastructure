


def main():
    data = [2,6,3,1,8,6,0,4]
    ret  = []
    for i in range(0,len(data)):
        for j in range(i+1, len(data)):
            if data[j] < data[i]:
                tmp = data[i]
                data[i] = data[j]
                data[j] = tmp

    print(data)



if __name__ == "__main__":
    main()