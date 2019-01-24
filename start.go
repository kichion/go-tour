package main

import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walkSub func(t *tree.Tree)
	walkSub = func(t *tree.Tree) {
        if t != nil {
            walkSub(t.Left)
            ch <- t.Value
            walkSub(t.Right)
        }
    }
	walkSub(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
        println(v)
    }
}