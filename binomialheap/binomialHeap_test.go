package binomialheap

import (
//    "fmt"
    "testing"
    . "priorityqueue"
//    . "time"
    . "priorityqueue/data"
)


func CreatePQElements(elements []Element) PriorityQueue {
    h := MakeHeap()
    for i:=0;i<len(elements);i++ {
        h.Insert(elements[i])
    }
    return h
}

var pqSamplesInsert1 = [][]Element{
    []Element{2,3,4},
    []Element{11,13,14,15,18},
    []Element{2,3},
    []Element{},
    []Element{},
}

var pqSamplesInsert2 = [][]Element{
    []Element{1,5,6},
    []Element{3,8,4,9},
    []Element{},
    []Element{3,8,4,9},
    []Element{},
}

var pqSamplesInsertOp = [][]Element{
    []Element{4,6,1,2,3,5},
    []Element{18,3,11,14,15,13,4,9,8},
    []Element{2,3},
    []Element{3,4,9,8},
    []Element{},
}

type test_Insert struct {
    e1 []Element
    e2 []Element
    expected []Element

}

var test_pairs_Insert = []test_Insert{
    {pqSamplesInsert1[0],pqSamplesInsert2[0],pqSamplesInsertOp[0]},
    {pqSamplesInsert1[1],pqSamplesInsert2[1],pqSamplesInsertOp[1]},
    {pqSamplesInsert1[2],pqSamplesInsert2[2],pqSamplesInsertOp[2]},
    {pqSamplesInsert1[3],pqSamplesInsert2[3],pqSamplesInsertOp[3]},
    {pqSamplesInsert1[3],pqSamplesInsert2[4],pqSamplesInsertOp[4]},
}

func TestUnion(t *testing.T) {
    for _,pair := range test_pairs_Insert {
        var pqc1, pqc2 PriorityQueue
        pqc1 = CreatePQElements(pair.e1)
        pqc2 = CreatePQElements(pair.e2)
        pq := Union(pqc1.(*BinomialHeap),pqc2.(*BinomialHeap))
        if !Equal(pq, pair.expected) {
            t.Error("For ",pair.e1, " & ", pair.e2," Expected ",pair.expected," got ", pq)
        }
    }
}

func Equal(pq1 PriorityQueue, e []Element) bool {
    pq1Iter := pq1.GetIterator()
    if pq1Iter == nil && len(e) == 0 {
        return true
    } else if pq1Iter == nil || len(e) == 0 {
        return false
    }
    var hasNext1, hasNext2 bool
    var e1 Element
    for i:=0;;i++ {
        e1,hasNext1 = pq1Iter()
        if i < len(e) - 1 {
            hasNext2 = true
        } else {
            hasNext2 = false
        }
//        fmt.Println("e1,e2::", e1, e[i])
        if e1 != e[i] || hasNext1 != hasNext2 {
            return false
        }
        if hasNext1 == false && hasNext2 == false {
            break
        }
    }
    return true
}


