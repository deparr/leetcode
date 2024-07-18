type Trie struct {
    children [26]*Trie
    valid bool
}


func Constructor() Trie {
    return Trie{}
}


func (this *Trie) Insert(word string)  {
    cur := this
    for _, c := range word {
        next := cur.children[c - 'a']
        if next == nil {
            new := Constructor()
            cur.children[c - 'a'] = &new
            next = &new
        }
        cur = next
    }

    cur.valid = true
}


func (this *Trie) Search(word string) bool {
    cur := this
    for _, c := range word {
        cur = cur.children[c - 'a']
        if cur == nil {
            return false
        }
    }

    return cur.valid
}


func (this *Trie) StartsWith(prefix string) bool {
    fmt.Println("prefixing", prefix)
    cur := this
    for _, c := range prefix {
        cur = cur.children[c - 'a']
        if cur == nil {
            return false
        }
    }

    if cur.valid {
        return true
    }

    for i, t := range cur.children {
        if t != nil {
            if t.valid || this.StartsWith(prefix + string('a' + i)) {
                return true
            }
        }
    }

    return false
}
