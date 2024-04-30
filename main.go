package main

import (
    "fmt"
    "math/rand"
    "sync"
)

type SafeMap struct {
    sync.Mutex
    map1 map[int]int
}

func (x *SafeMap) read(k int) (int, bool) {
    x.Lock()
    defer x.Unlock()
    val, ok := x.map1[k]
    return val, ok
}

func (x *SafeMap) write(p, v int) {
    x.Lock()
    defer x.Unlock()
    x.map1[p] = v
}

func (x *SafeMap) delete(n int) {
    x.Lock()
    defer x.Unlock()
    delete(x.map1, n)
}

func main() {
    var wg sync.WaitGroup
    x := &SafeMap{map1: make(map[int]int)}

    for j := 0; j < 10; j++ {
        wg.Add(1)
        go func(j int) {
            defer wg.Done()
            x.write(j, rand.Intn(100))
        }(j)
    }

    for k := 1; k < 11; k++ {
        wg.Add(1)
        go func(k int) {
            defer wg.Done()
            val, ok := x.read(k)
            if ok {
                fmt.Printf("%d, %d\n", k, val)
            } else {
                fmt.Printf("Key: %d, Value: mavjud emas\n", k)
            }
        }(k)
    }

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            x.delete(i)
        }(i)
    }

    wg.Wait()
}