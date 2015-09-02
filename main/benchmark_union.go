package main

import (
    "fmt"
    . "binomialheap"
    . "priorityqueue"
    . "time"
    . "priorityqueue/data"
)

func CreatePQElements(elements []Element) PriorityQueue {
    h := MakeHeap()
    for i:=0;i<len(elements);i++ {
        h.Insert(elements[i])
    }
    return h
}

func BenchmarkUnion1() {
    pq1 := CreatePQElements([]Element{4})
    pq2 := CreatePQElements([]Element{3})
    var ct Time
    var d,avg Duration
    for n:=0;n<10000000;n++ {
        ct = Now()
        pq1 = Union(pq1.(*BinomialHeap),pq2.(*BinomialHeap))
        d = Since(ct)
        avg += d;
        pq1 = CreatePQElements([]Element{4})
        pq2 = CreatePQElements([]Element{3})
    }
    fmt.Println("BenchmarkUnion1\t10000000\t",avg/10000000,"/op")
}

func BenchmarkUnion3() {
    pq1 := CreatePQElements([]Element{2,5,4})
    pq2 := CreatePQElements([]Element{3,7,8})
    var ct Time
    var d,avg Duration
    for n:=0;n<10000000;n++ {
        ct = Now()
        pq1 = Union(pq1.(*BinomialHeap),pq2.(*BinomialHeap))
        d = Since(ct)
        avg += d;
        pq1 = CreatePQElements([]Element{2,5,4})
        pq2 = CreatePQElements([]Element{3,7,8})
    }
    fmt.Println("BenchmarkUnion3\t10000000\t",avg/10000000,"/op")
}

func BenchmarkUnion5() {
    pq1 := CreatePQElements([]Element{3,4,5,7,9})
    pq2 := CreatePQElements([]Element{13,45,32,34,56})
    var ct Time
    var d,avg Duration
    for n:=0;n<10000000;n++ {
        ct = Now()
        pq1 = Union(pq1.(*BinomialHeap),pq2.(*BinomialHeap))
        d = Since(ct)
        avg += d;
        pq1 = CreatePQElements([]Element{3,4,5,7,9})
        pq2 = CreatePQElements([]Element{13,45,32,34,56})
    }
    fmt.Println("BenchmarkUnion5\t10000000\t",avg/10000000,"/op")
}

func BenchmarkUnion1M() {
    pq1 := CreatePQElements([]Element{})
    for i:=2;i<10000000;i++ {
         pq1.Insert(Element(i))
    }
    pq2 := CreatePQElements([]Element{1})
    var ct Time
    var d,avg Duration
    for n:=0;n<10;n++ {
        ct = Now()
        pq1 = Union(pq1.(*BinomialHeap),pq2.(*BinomialHeap))
        d = Since(ct)
        avg += d;
        pq1 = CreatePQElements([]Element{})
        for i:=0;i<10000000;i++ {
             pq1.Insert(Element(i))
        }
        pq2 = CreatePQElements([]Element{1})
    }
    fmt.Println("BenchmarkUnion5\t10\t",avg/10000000,"/op")
}


func main() {
     BenchmarkUnion1()
     BenchmarkUnion3()
     BenchmarkUnion5()
//     BenchmarkUnion1M()
}
