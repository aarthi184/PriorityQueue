package PQ_Creator

import (
    . "PriorityQueue"
    . "PriorityQueue/Data"
    "BinaryHeap"
    "BinomialHeap"
    "FibonacciHeap"
)


func CreatePQFromBinaryHeap() PriorityQueue {
    h := BinaryHeap.MakeHeap()
    return h
}

func CreatePQFromBinomialHeap() PriorityQueue {
    h := BinomialHeap.MakeHeap()
    return h
}

func CreatePQFromFibonacciHeap() PriorityQueue {
    h := FibonacciHeap.MakeHeap()
    return h
}


func CreatePQElements(elements []Element) PriorityQueue {
    h := CreatePQFromBinaryHeap()
    for i:=0;i<len(elements);i++ {
        h.Insert(elements[i])
    }
    return h
}

