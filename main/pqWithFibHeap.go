package main

import (
     "fmt"
     . "priorityqueue"
     . "creator"
)

func main () {
     fmt.Println("Hello!!")
     var pq PriorityQueue
     pq = CreatePQFromFibonacciHeap()
     fmt.Println("Hello!!")
     if pq.IsEmpty() {
          fmt.Println("Heap is empty")
     }
     pq.Display()
     pq.Insert(6)
     pq.Display()
     pq.Insert(7)
     pq.Display()
     pq.Insert(4)
     pq.Display()
     pq.Insert(8)
     pq.Display()
     pq.Insert(2)
     pq.Display()
     fmt.Println(pq.DeleteMin())
     pq.Display()
     pq.Insert(9)
     pq.Display()
     pq.Insert(2)
     pq.Display()
     pq.Insert(5)
     pq.Display()
     pq.Insert(1)
     pq.Display()
     pq.Insert(13)
     pq.Display()
     pq.Insert(15)
     pq.Display()
     pq.Insert(25)
     pq.Display()
     fmt.Println(pq.DeleteMin())
     pq.Display()
}
