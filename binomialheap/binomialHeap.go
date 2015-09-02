package binomialheap

import (
    . "priorityqueue/data"
    "fmt"
)


//type Element int

//type DataIter func() (Element, bool)


type BinomialHeap struct {
    key Element
    order int
    leftChild *BinomialHeap
    rightSibling *BinomialHeap
    parent *BinomialHeap
}

func MakeHeap() *BinomialHeap {
    return &BinomialHeap{}
}

func (b *BinomialHeap) Insert(e Element) {
    newHeap := MakeHeap()
    newHeap.rightSibling = MakeHeap()
    newHeap.rightSibling.key = e
    b = Union(b,newHeap)
    b.updateMin()
}

func (b *BinomialHeap) updateMin() {
    b.leftChild = b.findMinNode()
}

func (b BinomialHeap) FindMin() Element {
    minNode := b.leftChild
    if minNode != nil {
        return minNode.key
    }
    return -1
}

func (b BinomialHeap) findMinNode() *BinomialHeap {
    var min *BinomialHeap
    if b.rightSibling == nil {
        return nil
    }
    min = b.rightSibling
    next := b.rightSibling.rightSibling
    for next != nil {
        if next.key < min.key {
             min = next
        }
        next = next.rightSibling
    }
    return min
}

func (b *BinomialHeap) DecreaseKey(oldKey Element, newKey Element) bool {
    if (newKey >= oldKey) {
        return false
    }
    bNode := b.FindKey(oldKey)
    if bNode == nil {
        return false
    }
    bNode.key = newKey
    for bNode.parent != nil {
        if bNode.parent.key > bNode.key {
            swap(bNode, bNode.parent)
        } else {
            b.updateMin()
            return true
        }
        bNode = bNode.parent
    }
    b.updateMin()
    return true
}

func swap(x *BinomialHeap, y *BinomialHeap) {
    t := x.key
    x.key = y.key
    y.key = t
}

func (b *BinomialHeap) DeleteMin() Element {
    return b.ExtractMin()
}

func (b *BinomialHeap) findPrevOfMinNode() *BinomialHeap {
    var next, prev, prevMin *BinomialHeap
    if b.rightSibling == nil {
        return nil
    }
    prevMin = b
    prev = b.rightSibling
    next = b.rightSibling.rightSibling
    for next != nil {
        if next.key < prevMin.rightSibling.key {
             prevMin = prev
        }
        next = next.rightSibling
        prev = prev.rightSibling
    }
    return prevMin
}

func (b *BinomialHeap) ExtractMin() Element {
    var minNode, prevNode, newHeap *BinomialHeap
    prevNode = b.findPrevOfMinNode()
    minNode = prevNode.rightSibling
    prevNode.rightSibling = minNode.rightSibling
    minNode.rightSibling = nil
    if minNode.leftChild != nil {
        newHeap = minNode.leftChild
        newHeap = newHeap.reverseList()
        b = Union(b,newHeap)
    }
    b.updateMin()
    return minNode.key
}

// Reverses the nodes in the linked list
// Returns a heap (with header)
func (b *BinomialHeap) reverseList() *BinomialHeap {
    var head, cur *BinomialHeap
    head = MakeHeap()
    cur = b
    for cur != nil {
        t := head.rightSibling
        head.rightSibling = cur
        cur = cur.rightSibling
        head.rightSibling.parent = nil
        head.rightSibling.rightSibling = t
    }
    return head
}

func (b *BinomialHeap) Delete(key Element) bool {
    status := b.DecreaseKey(key, -1)
    if status == false {
        return false
    }
    e := b.ExtractMin()
    if (e == -1) {
        return true
    }
    return false
}

func (b *BinomialHeap) FindKey(key Element) *BinomialHeap {
    if b.rightSibling == nil {
        return nil
    }
    e := b.rightSibling.returnElement(key)
    return e
}


func (b *BinomialHeap) returnElement(key Element) *BinomialHeap {
    next := b
    var e *BinomialHeap
    e = nil
    for next != nil {
        if key < next.key {
             next = next.rightSibling
        } else if key == next.key {
            return next
        } else {
            if (next.leftChild != nil) {
                e = next.leftChild.returnElement(key)
                if e != nil {
                     return e
                }
            }
            next = next.rightSibling
        }
    }
    return nil
}

func (b *BinomialHeap) IsEmpty() bool {
    if b == nil || b.rightSibling == nil{
        return true
    }
    return false
}

func (b BinomialHeap) Display() {
    if b.rightSibling == nil {
        fmt.Println("No elements. Empty structure")
        return
    }
    b.rightSibling.displayKey(0)
}

func (b BinomialHeap) displayKey(level int) {
    next := &b
    for next != nil {
        if (level > 0) {
            fmt.Print("Child:");
        }
        fmt.Println(next.key, "Order", next.order, "Level", level);
        if (next.leftChild != nil) {
            next.leftChild.displayKey(level + 1)
        }
        next = next.rightSibling
    }
}

func Union(b *BinomialHeap, b2 *BinomialHeap) *BinomialHeap {
    var next *BinomialHeap
    var cur1, cur2 *BinomialHeap
    next = b
    cur1 = b.rightSibling
    cur2 = b2.rightSibling
    //Simple merge of the 2 binomial heaps
    //in non-decreasing order
    for cur1 != nil && cur2 != nil {
        if cur1.order <= cur2.order {
            next.rightSibling = cur1
            cur1 = cur1.rightSibling
        } else {
            next.rightSibling = cur2
            cur2 = cur2.rightSibling
        }
        if next.rightSibling == nil {
            break
        }
        next = next.rightSibling
    }
    if cur1 != nil {
        next.rightSibling = cur1
    }
    if cur2 != nil {
        next.rightSibling = cur2
    }
    //Combining binomial trees of same order
    var x, nextX, nextNextX, prevX *BinomialHeap
    prevX = nil
    x = b.rightSibling
    if x != nil {
        nextX = x.rightSibling
    } else {
        nextX = nil
    }
    nextNextX = nil
    for nextX != nil {
        if (nextX.rightSibling != nil) {
            nextNextX = nextX.rightSibling
        }
        if ((x.order != nextX.order) || (nextNextX != nil && x.order == nextNextX.order)) {
            prevX = x
            x = nextX
        } else if (x.key <= nextX.key) { //If 2 orders are equal
            x.rightSibling = nextNextX
            BinomialLink(x,nextX)
        } else if (prevX != nil) {
            prevX.rightSibling = nextX
            BinomialLink(nextX,x)
            x = prevX.rightSibling
        } else {
            b.rightSibling = nextX
            BinomialLink(nextX,x)
            x = b.rightSibling
        }
        nextNextX = nil
        nextX = x.rightSibling
    }
    b.updateMin()
    return b
}

func BinomialLink(x *BinomialHeap, y *BinomialHeap) {
    x.order +=1
    y.rightSibling = x.leftChild
    x.leftChild = y
    y.parent = x
}

func  (b *BinomialHeap) GetIterator() DataIter {
    if b.rightSibling == nil {
        return nil
    }
    cur := b.rightSibling
    return func() (e Element,hasNext bool) {
        e = cur.key
        if cur.leftChild != nil {
            cur = cur.leftChild
            return e, true
        } else if cur.parent == nil && cur.rightSibling == nil {
            return e,false
        } else if cur.rightSibling != nil {
            cur = cur.rightSibling
            return e, true
        } else if cur.parent != nil {
            cur = cur.parent.rightSibling
            if cur != nil {
                return e,true
            } else {
                return e, false
            }
        }
        return -1,false
    }
}



func (b BinomialHeap) GetSize() int {
    if b.rightSibling == nil {
        return 0
    }
    return b.rightSibling.calculateSize()
}

func (b BinomialHeap) calculateSize() int {
    next := &b
    size := 0
    for next != nil {
        size++
        if (next.leftChild != nil) {
            size += next.leftChild.calculateSize()
        }
        next = next.rightSibling
    }
     return size
}
