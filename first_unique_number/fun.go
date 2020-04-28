package fun

type Node struct {
    Val        int
    Prev, Next *Node
}

type FU struct {
    store      []int
    head, tail *Node
    dupSet     map[int]int
    curUniqSet map[int]*Node
}

func constructor(nums []int) *FU {
    set := make(map[int]int)
    curSet := make(map[int]*Node)
    for i := range nums {
        set[nums[i]]++
    }
    head, tail := &Node{Val: -1}, &Node{Val: -1}
    head.Next, tail.Prev = tail, head
    cur := head
    for i := range nums {
        if set[nums[i]] == 1 {
            node := &Node{Val: nums[i]}
            tmp := cur.Next
            cur.Next = node
            node.Prev = cur
            node.Next = tmp
            tmp.Prev = node
            curSet[nums[i]] = node
            cur = cur.Next
        }
    }
    return &FU{
        store:  nums,
        head:   head,
        tail:   tail,
        dupSet: set,
    }
}

func (fu FU) showFirstUnique() int {
    if fu.head.Next != fu.tail {
        return fu.head.Next.Val
    }
    return -1
}

// add time complexity O(N)
func (fu *FU) add(val int) {
    fu.store = append(fu.store, val)
    fu.dupSet[val]++
    count := fu.dupSet[val]
    if count > 2 {
        return
    } else if count == 2 {
        node := fu.curUniqSet[val]
        node.Prev.Next = node.Next
        node.Next.Prev = node.Prev
        delete(fu.curUniqSet, val)
    } else {
        node := &Node{Val: val, Prev: fu.tail.Prev, Next: fu.tail}
        fu.tail.Prev.Next = node
        fu.tail.Prev = node
        fu.curUniqSet[val] = node
    }
    return
}
