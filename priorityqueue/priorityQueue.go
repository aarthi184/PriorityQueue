package PriorityQueue

import (
    . "PriorityQueue/Data"
)


type PriorityQueue interface {
    DeleteMin() Element
    FindMin() Element
    Insert(Element)
    IsEmpty() bool
    Display()
    GetSize() int
    GetIterator() DataIter
}

