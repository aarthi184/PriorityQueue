package main

import (
    "fmt"
    . "BinomialHeap"
    . "PriorityQueue/Data"
)

func main() {
//    var b PriorityQueue
    var b *BinomialHeap
    b = MakeHeap()
    fmt.Println("Delete 20",b.Delete(20))
    fmt.Println("SIZE:",b.GetSize())
    b.Insert(35)
    fmt.Println("Displaying the heap")
    b.Display()
    fmt.Println("SIZE:",b.GetSize())
    b.Insert(62)
    fmt.Println("Displaying the heap")
    b.Display()
    fmt.Println("SIZE:",b.GetSize())
    b.Insert(22)
    fmt.Println("Displaying the heap")
    b.Display()
    fmt.Println("Minimum:", b.FindMin())
    b.Insert(89)
    fmt.Println("Displaying the heap")
    b.Display()
    fmt.Println("Minimum:", b.FindMin())
    b.Insert(31)
    fmt.Println("Displaying the heap")
    b.Display()
    fmt.Println("Minimum:", b.FindMin())
    b.Insert(20)
    fmt.Println("Displaying the heap")
    b.Display()
    fmt.Println("Minimum:", b.FindMin())
    b.Insert(26)
    fmt.Println("Displaying the heap")
    b.Display()
    b.Insert(5)
    fmt.Println("Displaying the heap")
    b.Display()
    fmt.Println("Minimum:", b.FindMin())
    b.Insert(18)
    fmt.Println("Displaying the heap")
    b.Display()
    fmt.Println("SIZE:",b.GetSize())

    fmt.Println("Find 20:", b.FindKey(20))
    fmt.Println("Find 89:", b.FindKey(89))
    fmt.Println("Find 62:", b.FindKey(62))
    fmt.Println("Find 1:", b.FindKey(1))
    fmt.Println("Find 90:", b.FindKey(90))

    b.DecreaseKey(89,3)
    fmt.Println("Displaying the heap")
    b.Display()
    fmt.Println("Find 1:", b.FindKey(1))
    fmt.Println("Minimum:", b.FindMin())
    fmt.Println("SIZE:",b.GetSize())

    fmt.Println("Min Extracted",b.ExtractMin())
    fmt.Println("Displaying the heap")
    b.Display()
    fmt.Println("Find 1:", b.FindKey(1))
    fmt.Println("Minimum:", b.FindMin())
    fmt.Println("SIZE:",b.GetSize())

    fmt.Println("Delete 20",b.Delete(20))
    fmt.Println("Displaying the heap")
    b.Display()
    fmt.Println("SIZE:",b.GetSize())
    b.Insert(89)
    b.Insert(1)
    b.Insert(2)
    fmt.Println("Displaying the heap")
    b.Display()
    fmt.Println("Delete 31",b.Delete(31))
    fmt.Println("SIZE:",b.GetSize())
    fmt.Println("Displaying the heap")
    b.Display()


    var i Element
    hasNext := true
    pq := b.GetIterator()
    for hasNext == true {
        i,hasNext = pq()
        fmt.Println(i)
    }
    fmt.Println("Displaying the heap")
    b.Display()
}
