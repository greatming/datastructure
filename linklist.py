
class Node:
    def __init__(self, initData = ""):
        self.__data = initData
        self.__next = None


    def setNext(self, newNext):
        self.__next = newNext

    def setData(self, data):
        self.__data = data

    def getData(self):
        return self.__data

    def getNext(self):
        return self.__next



class LinkList:
    def __init__(self):
        self.head = Node(None)
        self.head.setNext(self.head)

    def add(self, data):
        node = Node(data)
        node.setNext(self.head.getNext())
        self.head.setNext(node)

    def remove(self, data):
        cur = self.head.getNext()

        while cur != self.head:
            if data == cur.getNext().getData():
                cur.setNext(cur.getNext().getNext())
                break
            cur = cur.getNext()




    def getItemIndex(self,item):
        cur = self.head.getNext()
        index = 0
        isFind = False

        while cur != self.head:
            if item == cur.getData():
                isFind = True
                break
            index = index + 1
            cur = cur.getNext()

        if isFind:
            return index
        return -1



    def getList(self):
        restlist = []
        cur = self.head.getNext()
        while cur != self.head:
            restlist.append(cur.getData())
            cur = cur.getNext()
        return restlist


    def reverse(self):
        cur = self.head.getNext()
        prevNode = self.head

        while cur != self.head:
            nextNode = cur.getNext()
            cur.setNext(prevNode)
            prevNode = cur
            cur = nextNode
        self.head.setNext(prevNode)



def main():
    list = LinkList()
    list.add(1)
    list.add(2)
    list.add(3)
    list.add(4)
    list.add(5)
    retlist = list.getList()
    print(retlist)

    # index = list.getItemIndex(1)
    # print(index)
    #
    # list.remove(3)
    #
    # retlist = list.getList()
    # print(retlist)

    list.reverse()
    retlist = list.getList()
    print(retlist)


if __name__ == "__main__":
    main()
