package _622_Design_Circular_Queue

import (
    "fmt"
    "testing"
)

func TestSolution(t *testing.T) {
    obj := Constructor(3)
    fmt.Println(obj.EnQueue(1))
    fmt.Println(obj.EnQueue(2))
    fmt.Println(obj.EnQueue(3))
    fmt.Println(obj.EnQueue(4))
    fmt.Println(obj.Rear())
    fmt.Println(obj.IsFull())
    fmt.Println(obj.DeQueue())
    fmt.Println(obj.EnQueue(4))
    fmt.Println(obj.Rear())

    if false {
        t.Errorf("expected")
    }
}
