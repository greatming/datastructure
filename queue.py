
class Queue:
    def __init__(self, n):
        self.__head = 0
        self.__tail = 0
        self.__len = n
        self.__queue = [None] * n

    def push(self, item):
        if self.__tail == self.__len:
            if self.__head == 0:
                return False
            for i in range(self.__head, self.__tail):
                self.__queue[i - self.__head] = self.__queue[i]

            self.__tail -= self.__head
            self.__head = 0


        self.__queue[self.__tail] = item
        self.__tail += 1
        return True


    def push2(self, item):
        #循环队列队列已满的条件(tail + 1) % n  = head
        if ((self.__tail + 1) % self.__len) == self.__head:
            return False

        self.__queue[self.__tail] = item
        self.__tail = (self.__tail + 1) % self.__len
        return True


    def pop(self):
        if self.__head == self.__tail:
            return False
        ret = self.__queue[self.__head]
        self.__head += 1
        return ret



def main():
    que = Queue(5)
    que.push2(1)
    que.push2(2)
    que.push2(3)
    que.push2(4)
    ret = que.push2(5)
    print(ret)
    ret = que.pop()
    print(ret)
    ret = que.push2(6)
    print(ret)
    ret = que.pop()
    print(ret)


if __name__ == "__main__":
    main()