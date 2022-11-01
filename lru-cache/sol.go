type LRUCache struct {
    head *CacheNode
    tail *CacheNode
    
    pointers map[int]*CacheNode
    
    capacity int
    size int
}

type CacheNode struct {
    Key int
    Val int
    Next *CacheNode
    Prev *CacheNode
}


func Constructor(capacity int) LRUCache {
    nhead := CacheNode{Key: -1, Val: -1, Next: nil, Prev: nil}
    return LRUCache{head: &nhead, tail: &nhead, pointers: map[int]*CacheNode{}, capacity: capacity, size: 0}
}


func (this *LRUCache) Get(key int) int {

    node, pres := this.pointers[key]
    if !pres {
        return -1
    }
    
    if node != this.head {

        if node == this.tail {
            this.tail = node.Prev
        }

        if node.Next != nil {
            node.Next.Prev = node.Prev
        }
        node.Prev.Next = node.Next
        


        node.Next = this.head
        this.head.Prev = node
        node.Prev = nil
        this.head = node   
    }


    return node.Val
}


func (this *LRUCache) Put(key int, value int)  {
    if this.size == 0 {
        this.head.Key = key
        this.head.Val = value
        this.size++
        this.pointers[key] = this.head
        return
    }

    node, pres := this.pointers[key]
    if pres {
        node.Val = value
        if node == this.head {
            return
        }

        node.Prev.Next = node.Next
        if node == this.tail {
            this.tail = this.tail.Prev
        } else {
            node.Next.Prev = node.Prev
        }
        node.Prev = nil
        this.head.Prev = node
        node.Next = this.head
        this.head = node
        return
    }

    new_node := &CacheNode{Key: key, Val: value, Next: this.head, Prev: nil}
    this.head.Prev = new_node
    this.head = new_node

    this.size++
    if this.size > this.capacity  {
        this.size--

        delete(this.pointers, this.tail.Key)
        this.tail = this.tail.Prev
        this.tail.Next = nil
    }
    this.pointers[key] = new_node
}
