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

func (fh *FibonacciHeap) Delete() {

}

func (fh FibonacciHeap) FindKey(e Element) *fHeapNode {
    if fh.min != nil {
        return fh.min.findNode(e)
    }
    return nil
}

func (fNode *fHeapNode) findNode(e Element) *fHeapNode {
    if fNode.key == e {
        return fNode
    }
    next := fNode.left
    for next != nil {
        if next.key == e {
            return next
        }
        if next.child != nil {
            node := next.child.findNode(e)
            if node != nil {
                 return node
            }
        }
        next = next.left
    }
    next = fNode.right
    for next != nil {
        if next.key == e {
            return next
        }
        if next.child != nil {
            node := next.child.findNode(e)
            if node != nil {
                 return node
            }
        }
        next = next.right
    }
    return nil
}

func (fNode *fHeapNode) cutOffNode() {
    if fNode.parent != nil {
        fNode.parent.child = fNode.right
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
        fh.addToRootList(nextChild)
        fh.updateMin(nextChild)
        nextChild = nextChild.right
    }
    if single == true {
        return minNode.key
    }
    fh.mergeTrees()
    return minNode.key
}

func (fh *FibonacciHeap) mergeTrees() {
    var array []*fHeapNode
    var parent, child *fHeapNode
    next := fh.min
    right := fh.min.right
    for next != nil {
        if next.order >= len(array) {
            count := next.order + 1 - len(array)
            for i:=0;i<count;i++ {
                array = append(array,nil)
            }
        }
        if array[next.order] == nil {
            array[next.order] = next
            next = next.left
        } else {
            if array[next.order].key < next.key {
                child = next
                parent = array[next.order]
                next = next.left
                makeChild(parent, child)
            } else {
                parent = next
                child = array[next.order]
                next = next.left
                makeChild(parent, child)
           }
        }
    }
    next = right
    for next != nil {
        if next.order >= len(array) {
            count := next.order + 1 - len(array)
            for i:=0;i<count;i++ {
                array = append(array,nil)
            }
        }
        if array[next.order] == nil {
            array[next.order] = next
            next = next.right
        } else {
            if array[next.order].key < next.key {
                child = next
                parent = array[next.order]
                next = next.right
                makeChild(parent, child)
            } else {
                parent = next
                child = array[next.order]
                next = next.right
                makeChild(parent, child)
           }
        }
    }
//        fmt.Println(array)
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

    return -1
}

func (fh FibonacciHeap) Display() {
    if fh.min == nil {
         fmt.Println("No elements. Empty Structure!")
         return
    }
    fmt.Println("Displaying the heap:")
    next := fh.min
    for next != nil {
        fmt.Println("Key:", next.key, "Marked:", next.marked)
        fh.min.displayKeys()
        next = next.left
    }
    next = fh.min.right
    fmt.Println("Right")
    for next != nil {
        fmt.Println("Key:", next.key, "Marked:", next.marked)
        fh.min.displayKeys()
        next = next.right
    }
}

func (fNode fHeapNode) displayKeys() {
    nextChild := fNode.child
    for nextChild != nil {
        fmt.Println("Children of ", fNode.key, "::")
        fmt.Println("Key:", nextChild.key, "Marked:", nextChild.marked)
        nextChild.displayKeys()
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
    return func () (Element,bool) {
         return -1,false
    }
}

