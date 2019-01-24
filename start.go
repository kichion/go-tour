package main

import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walk_sub func(t *tree.Tree)
	walk_sub = func(t *tree.Tree) {
        if t != nil {
            walk_sub(t.Left)
            ch <- t.Value
            walk_sub(t.Right)
        }
    }
	walk_sub(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	treeCh1 := make(chan int)
    treeCh2 := make(chan int)
    go Walk(t1, treeCh1)
    go Walk(t2, treeCh2)
    for v1 := range treeCh1 {
        v2, ok := <- treeCh2
        if !ok || v1 != v2 {
            return false
        }
    }
    return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
        println(v)
    }
	println(Same(tree.New(1), tree.New(1)))
	println(Same(tree.New(1), tree.New(2)))
}
