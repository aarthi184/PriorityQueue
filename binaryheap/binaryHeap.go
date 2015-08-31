package binaryheap

import (
    "fmt"
    . "priorityqueue/data"
)

//type Element int

//type DataIter func() (Element, bool)

type Heap struct {
    array []Element
    size int
}

//Need to change bool to hasnext
func (h Heap) GetIterator() DataIter {
    if h.size <= 0 {
        return nil
    }
    i := 0
    return func() (e Element,hasNext bool) {
        e = h.array[i]
        i++
        if i < h.size {
            hasNext = true
        } else {
            hasNext = false
        }
        return e,hasNext
    }
}

func (h *Heap) GetSize() int {
    return h.size
}

func (h *Heap) GetElements() []Element {
    return h.array
}

func MakeHeap() *Heap {
    return &Heap {
        array: []Element{},
    }
}

func (h *Heap) Insert(x Element) {
//    h.array[h.size] = x
    h.array = append(h.array[:h.size], x)
    h.size++
    h.siftUp()
}

func (h *Heap) siftUp() {
    var i, parent int
    i = h.size - 1
    for i > 0 {
        parent = (i - 1) /2
        if h.array[i] < h.array[parent] {
            temp := h.array[i]
            h.array[i] = h.array[parent]
            h.array[parent] = temp
        } else {
            return
        }
        i = parent
    }
}


func (h *Heap) FindMin() Element {
    if h.size > 0 {
        return h.array[0];
    } else {
        return -1;
    }
}

func (h * Heap) IsEmpty() bool {
    if h.size == 0 {
        return true
    } else {
        return false
    }
}

func (h * Heap) DeleteMin() Element {
    min := h.array[0]
    h.array[0] = h.array[h.size - 1]
    h.size--
    h.siftDown()
    return min
}

func (h *Heap) siftDown() {
     var i, lChild, rChild int
    i = 0
    for i < h.size {
        lChild = (2 * i) + 1
        rChild = (2 * i) + 2
        if rChild < h.size && lChild < h.size {
            if h.array[i] < h.array[lChild] && h.array[i] < h.array[rChild] {
                return
            }
            if h.array[lChild] < h.array[rChild] {
                h.swap(i,lChild)
                i = lChild
            } else {
                h.swap(i,rChild)
                i = rChild
            }
        } else if rChild >= h.size && lChild >= h.size {
            return
        } else if rChild >= h.size {
            if h.array[i] > h.array[lChild]{
                h.swap(i,lChild)
                i = lChild
            } else {
                return
            }
        } else {
            if h.array[i] > h.array[rChild]{
                h.swap(i,rChild)
                i = rChild
            } else {
                return
            }
        }
    }
}

func (h *Heap) swap(x int, y int) {
            temp := h.array[x]
            h.array[x] = h.array[y]
            h.array[y] = temp
}

func (h *Heap) Display() {
    if h.IsEmpty() {
         fmt.Println("Heap is empty")
         return
    }
    fmt.Println(h.array[:h.size])
}

