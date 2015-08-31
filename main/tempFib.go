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
    fmt.Println("Insert 16,26,36")
     pq.Insert(16)
     pq.Insert(26)
     pq.Insert(36)
     pq.Display()

     fmt.Println("Extracting minimum",pq.ExtractMin())
     pq.Display()
     fmt.Println("Extracting minimum",pq.ExtractMin())
     pq.Display()
}
