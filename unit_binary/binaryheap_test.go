package test_pq

import (
    "fmt"
    "testing"
    . "priorityqueue"
    . "creator"
    . "time"
    . "priorityqueue/data"
)


var pqSamples = []PriorityQueue{
    CreatePQElements([]Element{2,3,4},"binaryheap"),
    CreatePQElements([]Element{1,3,4,5,8},"binaryheap"),
    CreatePQElements([]Element{2,3},"binaryheap"),
    CreatePQElements([]Element{7,8,10,30,39,40,50},"binaryheap"),
}

type test_pairs_delete struct {
    pq PriorityQueue
    min Element
}

var tests_delete = []test_pairs_delete{
    {pqSamples[0],2},
    {pqSamples[1],1},
    {pqSamples[2],2},
    {pqSamples[3],7},
}

func TestFindMin(t *testing.T) {
    for _,pair := range tests_delete {
        v := pair.pq.FindMin()
        if v!=pair.min {
            t.Error("For ",pair.pq," Expected ",pair.min," got ", v)
        }
    }
}

func TestDeleteMin(t *testing.T) {
    for _,pair := range tests_delete {
        v := pair.pq.DeleteMin()
        if v!=pair.min {
            t.Error("For ",pair.pq," Expected ",pair.min," got ", v)
        }
    }
}


var pqSamplesIsEmpty = []PriorityQueue{
    CreatePQElements([]Element{2},"binaryheap"),
    CreatePQElements([]Element{1,3},"binaryheap"),
    CreatePQElements([]Element{},"binaryheap"),
}

type test_IsEmpty struct {
    pq PriorityQueue
    result bool
}

var test_pairs_IsEmpty = []test_IsEmpty{
    {pqSamplesIsEmpty[0],false},
    {pqSamplesIsEmpty[1],false},
    {pqSamplesIsEmpty[2],true},
}

func TestIsEmpty(t *testing.T) {
    for _,pair := range test_pairs_IsEmpty {
        v := pair.pq.IsEmpty()
        if v!=pair.result {
            t.Error("For ",pair.pq," Expected ",pair.result," got ", v)
        }
    }
}

var pqSamplesInsert = []PriorityQueue{
    CreatePQElements([]Element{2,3,4},"binaryheap"),
    CreatePQElements([]Element{11,13,14,15,18},"binaryheap"),
    CreatePQElements([]Element{2,3},"binaryheap"),
    CreatePQElements([]Element{},"binaryheap"),
}

var pqSamplesInsertOp = []PriorityQueue{
    CreatePQElements([]Element{1,2,4,3},"binaryheap"),
    CreatePQElements([]Element{10,13,11,15,18,14},"binaryheap"),
    CreatePQElements([]Element{2,3,5},"binaryheap"),
    CreatePQElements([]Element{10},"binaryheap"),
}

type test_Insert struct {
    pq PriorityQueue
    x Element
    expectedPQ PriorityQueue

}

var test_pairs_Insert = []test_Insert{
    {pqSamplesInsert[0],1,pqSamplesInsertOp[0]},
    {pqSamplesInsert[1],10,pqSamplesInsertOp[1]},
    {pqSamplesInsert[2],5,pqSamplesInsertOp[2]},
    {pqSamplesInsert[3],10,pqSamplesInsertOp[3]},
}

func TestInsert(t *testing.T) {
    for _,pair := range test_pairs_Insert {
        pair.pq.Insert(pair.x)
        if !Equal(pair.pq, pair.expectedPQ) {
            t.Error("For ",pair.pq," Expected ",pair.expectedPQ," got ", pair.pq)
        }
    }
}

func Equal(pq1 PriorityQueue, pq2 PriorityQueue) bool {
    pq1Iter := pq1.GetIterator()
    pq2Iter := pq2.GetIterator()
    if pq1Iter == nil && pq2Iter == nil {
        return true
    } else if pq1Iter == nil || pq2Iter == nil {
        return false
    }
    var hasNext1, hasNext2 bool
    var e1,e2 Element
    for {
        e1,hasNext1 = pq1Iter()
        e2,hasNext2 = pq2Iter()
    //    fmt.Println("e1,e2::", e1, e2)
        if e1 != e2 || hasNext1 != hasNext2 {
            return false
        }
        if hasNext1 == false && hasNext2 == false {
            break
        }
    }
    return true
}

func BenchmarkInsert3(b *testing.B) {
    pq := CreatePQElements([]Element{2,3,4},"binaryheap")
    for n:=0;n<b.N;n++ {
        pq.Insert(5)
    }
}

func BenchmarkInsert6(b *testing.B) {
    pq := CreatePQElements([]Element{2,3,4,5,6,7},"binaryheap")
    for n:=0;n<b.N;n++ {
        pq.Insert(1)
    }
}

func BenchmarkInsert1M(b *testing.B) {
    pq := CreatePQElements([]Element{},"binaryheap")
    for i:=1;i <= 100; i++ {
         pq.Insert(Element(i))
    }
    for i:=102;i <= 10000000; i++ {
         pq.Insert(Element(i))
    }
//    pq.Display()
    for n:=0;n<b.N;n++ {
        pq.Insert(101)
    }
}

func BenchmarkIsEmpty1(b *testing.B) {
    pq := CreatePQElements([]Element{},"binaryheap")
    for n:=0;n<b.N;n++ {
        pq.IsEmpty()
    }
}

func BenchmarkIsEmpty2(b *testing.B) {
    pq := CreatePQElements([]Element{7,8,9,0},"binaryheap")
    for n:=0;n<b.N;n++ {
        pq.IsEmpty()
    }
}

func BenchmarkFindMin(b *testing.B) {
    pq := CreatePQElements([]Element{2,3,4,5,6},"binaryheap")
    for n:=0;n<b.N;n++ {
        pq.FindMin()
    }
}

func TestDeleteMinBenchmark100(t *testing.T) {
    pq := CreatePQElements([]Element{},"binaryheap")
    for i:=1;i <= 100; i++ {
         pq.Insert(Element(i))
    }
    var ct Time
    var d,avg Duration
    for n:=0;n<10000000;n++ {
        ct = Now()
        x := pq.DeleteMin()
        d = Since(ct)
        avg += d;
        pq.Insert(x)
    }
    fmt.Println("BenchmarkDeleteMin100\t10000000\t",avg/10000000,"/op")
}

func TestDeleteMinBenchmark1M(t *testing.T) {
    pq := CreatePQElements([]Element{},"binaryheap")
    for i:=1;i <= 10000000; i++ {
         pq.Insert(Element(i))
    }
    var ct Time
    var d,avg Duration
    for n:=0;n<10000000;n++ {
        ct = Now()
        x := pq.DeleteMin()
        d = Since(ct)
        avg += d;
        pq.Insert(x)
    }
    fmt.Println("BenchmarkDeleteMin1M\t10000000\t",avg/10000000,"/op")
}

