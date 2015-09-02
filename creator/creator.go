package creator

import (
    . "priorityqueue"
    . "priorityqueue/data"
    "binaryheap"
    "binomialheap"
    "fibonacciheap"
)


func CreatePQFromBinaryHeap() PriorityQueue {
    h := binaryheap.MakeHeap()
    return h
}

func CreatePQFromBinomialHeap() PriorityQueue {
    h := binomialheap.MakeHeap()
    return h
}

func CreatePQFromFibonacciHeap() PriorityQueue {
    h := fibonacciheap.MakeHeap()
    return h
}


func CreatePQElements(elements []Element, impl string) PriorityQueue {
    var h PriorityQueue
    if impl == "binomialheap" {
        h = CreatePQFromBinomialHeap()
    } else{
        h = CreatePQFromBinaryHeap()
    }
    for i:=0;i<len(elements);i++ {
        h.Insert(elements[i])
    }
    return h
}

