package main

import (
     "fmt"
     . "fibonacciheap"
     . "priorityqueue/data"
)

func main () {
     fmt.Println("Hello!!")
     pq := MakeHeap()


     pq.Insert(16)
     pq.Display()
     fmt.Println("Extracting minimum",pq.ExtractMin())
     pq.Display()
    fmt.Println("Insert 16,26,36,46")
     pq.Insert(16)
     pq.Insert(26)
     pq.Insert(36)
     pq.Insert(46)
     pq.Display()
     pq.Insert(56)
     pq.Display()

     fmt.Println("Extracting minimum",pq.ExtractMin())
     pq.Display()
     pq.Insert(16)
     pq.Display()
     pq.Insert(15)
     pq.Display()
     pq.Insert(14)
     pq.Insert(41)
     pq.Insert(31)
     pq.Insert(21)
     pq.Display()
     fmt.Println("Extracting minimum",pq.ExtractMin())
     pq.Display()
     pq.Insert(81)
     pq.Display()
     fmt.Println("Decrease key 26 - 12",pq.DecreaseKey(26,12))
     pq.Display()
     fmt.Println("Decrease key22 - 11",pq.DecreaseKey(22,11))
     pq.Insert(10)
     pq.Insert(9)
     pq.Display()
     fmt.Println("Decrease key 56-25",pq.DecreaseKey(56,25))
     pq.Display()
     fmt.Println("Decrease key 46-8",pq.DecreaseKey(46,8))
     pq.Display()
     fmt.Println("Decrease key 36-18",pq.Delete(36))
     pq.Insert(7)
     pq.Insert(6)
     pq.Insert(5)
     pq.Insert(4)
     pq.Insert(3)
     fmt.Println("Delete 5",pq.Delete(5))
     pq.Display()
    var i Element
    hasNext := true
    iter := pq.GetIterator()
    for hasNext == true {
        i,hasNext = iter()
        fmt.Println(i)
    }
}
