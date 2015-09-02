package fibonacciheap

import (
    . "priorityqueue/data"
    "fmt"
)

type FibonacciHeap struct {
    min *fHeapNode
    n int
}


type fHeapNode struct {
    key Element
    order int
    right *fHeapNode
    left *fHeapNode
    child *fHeapNode
    marked bool
    parent *fHeapNode
}

func (fh FibonacciHeap) FindMin() Element {
    if fh.min != nil {
        return fh.min.key
    }
    return -1
}

func MakeHeap() *FibonacciHeap {
    return &FibonacciHeap{}
}

func makeFHeapNode(e Element) *fHeapNode {
    return &fHeapNode{
        key: e,
    }
}

func (fh *FibonacciHeap) Insert(key Element) {
    node := makeFHeapNode(key)
    if fh.min == nil {
         fh.min = node
         fh.n++
         return
    }
    fh.addToRootList(node)
    fh.n++
}

func (fh *FibonacciHeap) addToRootList(node *fHeapNode) {
    node.left = nil
    node.right= nil
    node.parent= nil
    if fh.min == nil {
         fh.min = node
         return
    }
    // Inserting the node into the rootlist
    node.left = fh.min.left
//    fmt.Println("In insert node.l", node.left)
//    fmt.Println("In insert node.r", node.right)
//    fmt.Println("In insert node.key", node.key)
    node.right = fh.min
    fh.min.left = node
    if node.left != nil {
        node.left.right = node
    }
    // Moving the min pointer if necessary
    if fh.min.key > node.key {
        fh.min = node
    }
}

func (fh *FibonacciHeap) Delete(key Element) bool {
    status := fh.DecreaseKey(key, -1)
    if status == false {
        return false
    }
    fh.Display()
    e := fh.ExtractMin()
    if (e == -1) {
        return true
    }
    return false
}

func (fh FibonacciHeap) FindKey(e Element) *fHeapNode {
    if fh.min == nil {
        return nil
    }
    next := fh.min
    for next != nil {
        if next.key == e {
             return next
        }
        fNode := next.findNode(e)
        if fNode != nil {
            return fNode
        }
        next = next.left
    }
    next = fh.min.right
    for next != nil {
        if next.key == e {
             return next
        }
        fNode := next.findNode(e)
        if fNode != nil {
            return fNode
        }
        next = next.right
    }
    return nil
}

func (fNode *fHeapNode) findNode(e Element) *fHeapNode {
    nextChild := fNode.child
    for nextChild != nil {
        if nextChild.key == e {
             return nextChild
        }
        node := nextChild.findNode(e)
        if node != nil {
             return node
        }
        nextChild = nextChild.right
    }
    return nil
}

func (fNode *fHeapNode) cutOffNode() {
    if fNode.parent != nil {
        fNode.parent.child = fNode.right
        fNode.parent.order--
        if fNode.right != nil {
            fNode.right.left = nil
        }
    }
}

func (fh *FibonacciHeap) DecreaseKey(oldKey Element, newKey Element) bool {
    if (newKey >= oldKey) {
        return false
    }
    fNode := fh.FindKey(oldKey)
    if fNode == nil {
        return false
    }
    fNode.key = newKey
    // In case of root node update min value and return
    if fNode.parent == nil {
        fh.updateMin(fNode)
        return true
    }
    // Removing node from tree and marking the parent or cutting off the parent
    parent := fNode.parent
    if parent != nil && parent.key > fNode.key {
        for fNode != nil  {
            if fNode.parent == nil {
                return true
            }
            fNode.cutOffNode()
            fh.addToRootList(fNode)
            fh.updateMin(fNode)
            if parent.marked == false {
                parent.marked = true
                break
            } else {
                parent.marked = false
            }
            fNode = parent
            parent = fNode.parent
        }
    }
    return true
}

func (fh *FibonacciHeap) updateMin(fNode *fHeapNode) {
    if fh.min == nil {
        return
    }
    if fh.min.key > fNode.key {
        fh.min = fNode
    }
}

func (fh *FibonacciHeap) traverseAndUpdateMin() {
    if fh.min == nil {
        return
    }
    min := fh.min
    next := min.left
    for next != nil {
        if next.key <= fh.min.key {
            fh.min = next
        }
        next = next.left
    }
    next = min.right
    for next != nil {
        if next.key <= fh.min.key {
            fh.min = next
        }
        next = next.right
    }
}


func (fh *FibonacciHeap) ExtractMin() Element {
    var nextChild, minNode *fHeapNode
    if fh.min == nil {
        return -1
    }
    minNode = fh.min
    single := false
    if fh.min.left != nil && fh.min.right != nil {
        fh.min.left.right = fh.min.right
        fh.min.right.left = fh.min.left
        fh.min = fh.min.right
        fh.traverseAndUpdateMin()
    } else if fh.min.left != nil {
        fh.min.left.right = fh.min.right
        fh.min = fh.min.left
        fh.traverseAndUpdateMin()
    } else if fh.min.right != nil {
        fh.min.right.left = fh.min.left
        fh.min = fh.min.right
        fh.traverseAndUpdateMin()
    } else {
        single = true
        fh.min = nil
    }
    nextChild = minNode.child
    for nextChild != nil {
        child := nextChild
        nextChild = nextChild.right
        fh.addToRootList(child)
        fh.updateMin(child)
    }
    if single == true {
        return minNode.key
    }
    fmt.Println("In extractmin displaying")
    fh.Display()
    fh.mergeTrees()
    fh.n--
    return minNode.key
}

func (fh *FibonacciHeap) mergeTrees() {
    var orderArr []*fHeapNode
    var parent, child *fHeapNode
    next := fh.min
    right := fh.min.right
    for next != nil {
        leftNode := next.left
        if next.order >= len(orderArr) {
            count := next.order + 1 - len(orderArr)
            for i:=0;i<count;i++ {
                orderArr = append(orderArr,nil)
            }
        }
        for orderArr[next.order] != nil {
            oldOrder := next.order
            if orderArr[next.order].key < next.key {
                child = next
                parent = orderArr[next.order]
                makeChild(parent, child)
            } else {
                parent = next
                child = orderArr[next.order]
                makeChild(parent, child)
            }
            orderArr[oldOrder] = nil
            next = parent
            if next.order >= len(orderArr) {
                orderArr = append(orderArr,nil)
            }
        }
        orderArr[next.order] = next
        next = leftNode
    }
    next = right
    for next != nil {
        rightNode := next.right
        if next.order >= len(orderArr) {
            count := next.order + 1 - len(orderArr)
            for i:=0;i<count;i++ {
                orderArr = append(orderArr,nil)
            }
        }
        for orderArr[next.order] != nil {
            oldOrder := next.order
            if orderArr[next.order].key < next.key {
                child = next
                parent = orderArr[next.order]
                makeChild(parent, child)
            } else {
                parent = next
                child = orderArr[next.order]
                makeChild(parent, child)
            }
            orderArr[oldOrder] = nil
            next = parent
            if next.order >= len(orderArr) {
                orderArr = append(orderArr,nil)
            }
        }
        orderArr[next.order] = next
        next = rightNode
    }
    i := 0
//    for k:=0;k<len(orderArr);k++ {
//        fmt.Println("orderArr[", k, "] :", orderArr[k])
//        if orderArr[k] != nil {
//        fmt.Println("orderArr[", k, "].key :", orderArr[k].key)
//        }
//    }
    for orderArr[i] == nil {
        i++
    }
    orderArr[i].left = nil
    for i=0;i<len(orderArr); {
        j := i + 1
        if orderArr[i] != nil {
            for j< len(orderArr) && orderArr[j] == nil {
                j++
            }
                if j >= len(orderArr) {
                    orderArr[i].right = nil
                    break
                }
            orderArr[i].right = orderArr[j]
            orderArr[j].left = orderArr[i]
            orderArr[j].right = nil
 //       fmt.Println("orderArr[", i, "].key :", orderArr[i].key)
 //       fmt.Println("orderArr[", i, "].right.key :", orderArr[i].right.key)
 //       fmt.Println("orderArr[", j, "].key :", orderArr[j].key)
 //       fmt.Println("orderArr[", j, "].left.key :", orderArr[j].left.key)
//            orderArr[j] = nil
            orderArr[i] = nil
            i = j
        } else {
             i++
        }
    }
    fh.traverseAndUpdateMin()
}




// min pointer will not be updated here
func makeChild(toBeParent *fHeapNode, child *fHeapNode) {

    if child.left != nil {
        child.left.right = child.right
    }
    if child.right != nil {
        child.right.left = child.left
    }
    child.parent = toBeParent
    child.right = toBeParent.child
    if child.right != nil {
        child.right.left = child.right
    }
    child.left = nil
    toBeParent.child = child
    toBeParent.order++
}


//func (fh *FibonacciHeap) mergeTrees() {
//    if fNode.key == e {
//        return fNode
//    }
//    next := fNode.left
//    for next != nil {
//        if next.key == e {
//            return next
//        }
//        if next.child != nil {
//            node := next.child.findNode(e)
//            if node != nil {
//                 return node
//            }
//        }
//        next = next.left
//    }
//    next = fNode.right
//    for next != nil {
//        if next.key == e {
//            return next
//        }
//        if next.child != nil {
//            node := next.child.findNode(e)
//            if node != nil {
//                 return node
//            }
//        }
//        next = next.right
//    }
//    return nil

//}


func (fh *FibonacciHeap) DeleteMin() Element {
    return fh.ExtractMin()
}

func (fh FibonacciHeap) Display() {
    if fh.min == nil {
         fmt.Println("No elements. Empty Structure!")
         return
    }
    fmt.Println("Displaying the heap:")
    next := fh.min
    for next != nil {
        fmt.Println("Key:", next.key, "Marked:", next.marked, "Order:", next.order)
        next.displayChildren()
        next = next.left
    }
    next = fh.min.right
    fmt.Println("Right")
    for next != nil {
        fmt.Println("Key:", next.key, "Marked:", next.marked, "Order:", next.order)
        next.displayChildren()
        next = next.right
    }
}

func (fNode fHeapNode) displayChildren() {
    nextChild := fNode.child
    for nextChild != nil {
        fmt.Println("Children of ", fNode.key, "::")
        fmt.Println("Key:", nextChild.key, "Marked:", nextChild.marked, "Order:", nextChild.order)
        nextChild.displayChildren()
        fmt.Println("Back to parent");
        nextChild = nextChild.right
    }
}

func (fh FibonacciHeap) IsEmpty() bool {
    if fh.min == nil {
         return true
    }
    return false
}

func (fh FibonacciHeap) GetSize() int {
    return fh.n
}

func (fh FibonacciHeap) GetIterator() DataIter {
    if fh.min == nil {
        return nil
    }
    cur := fh.min
    dirRight := false
    return func () (e Element, hasNext bool) {
        e = cur.key
        if cur.child != nil {
            cur = cur.child
            return e, true
        }
        if cur.parent == nil {
            cur, hasNext = fh.moveInRootList(cur,&dirRight)
            return e, hasNext

        }
        if cur.right != nil && cur.parent != nil {
            cur = cur.right
            return e, false
        }
        if cur.right == nil && cur.parent != nil {
            cur = cur.parent
            for cur.right == nil && cur.parent != nil {
                cur = cur.parent
            }
            if cur.parent == nil {
                cur, hasNext = fh.moveInRootList(cur,&dirRight)
                return e, hasNext
            }
            if cur.right != nil {
                 cur = cur.right
                 return e, true
            }
        }
    return -1, false
    }
    return nil
}

func (fh FibonacciHeap) moveInRootList(node *fHeapNode, dirRight *bool) (cur *fHeapNode, hasNext bool) {
    cur = node
    if *dirRight == false {
        if cur.left != nil {
           cur = cur.left
            return cur, true
        } else if fh.min.right != nil {
            cur = fh.min.right
            *dirRight = true
            return cur, true
        } else {
            cur = nil
            return cur, false
        }
    } else {
        if cur.right != nil {
            cur = cur.right
            return cur, true
        } else {
            cur = nil
            return cur, false
        }
    }
}
