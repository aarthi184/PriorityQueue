package main

import (
     "fmt"
     . "FibonacciHeap"
)

func main () {
     fmt.Println("Hello!!")
     pq := MakeHeap()


     pq.Insert(16)
     pq.Display()
     fmt.Println("Extracting minimum",pq.ExtractMin())
     pq.Display()


     fmt.Println("Hello!!")
     if pq.IsEmpty() {
          fmt.Println("Heap is empty")
     }
     pq.Display()
     pq.Insert(16)
     pq.Display()
          fmt.Println("Display heap after decrease key with greater key", pq.DecreaseKey(16,61))
     pq.DecreaseKey(2,0)
          fmt.Println("Display heap after decrease key with invalid key")
     pq.Display()
     pq.DecreaseKey(16,15)
          fmt.Println("Display heap after decrease key")
     pq.Display()
     pq.Insert(88)
     pq.Display()
     pq.Insert(41)
     pq.Display()
     pq.DecreaseKey(88,45)
          fmt.Println("Display heap after decrease key")
     pq.Display()
     pq.Insert(31)
     pq.Display()
     pq.Insert(22)
     pq.Display()
//     fmt.Println(pq.DeleteMin())
     pq.Display()
     pq.Insert(19)
     pq.Display()
     pq.Insert(21)
     pq.Display()
     pq.Insert(71)
     pq.Display()
     pq.Insert(11)
     pq.Display()
     pq.Insert(53)
     pq.Display()
     pq.Insert(51)
     pq.Display()
     pq.Insert(25)
     pq.Display()
     fmt.Println("Extracting minimum",pq.ExtractMin())
     pq.Display()
     pq.DecreaseKey(13,8)
     pq.Display()
     pq.DecreaseKey(41,7)
     pq.Display()
     pq.DecreaseKey(15,6)
     pq.Display()
     fmt.Println("Extracting minimum",pq.ExtractMin())
     pq.Display()
}
