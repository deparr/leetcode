type Trie struct {
    TrieNode
}

type TrieNode struct {
    ch rune
    children [26]*TrieNode
    isEnd bool
}


func Constructor() Trie {
    return Trie{}
}


func (this *Trie) Insert(word string)  {
    cur := &this.TrieNode
    for _, c := range word {
        next := cur.children[c - 'a']
        if next == nil {
            new := &TrieNode{
                ch: c,
                children: [26]*TrieNode{},
                isEnd: false,
            }
            cur.children[c - 'a'] = new
            next = new
        }
        cur = next
    }

    cur.isEnd = true
}


func (this *Trie) Search(word string) bool {
    cur := &this.TrieNode
    for _, c := range word {
        cur = cur.children[c - 'a']
        if cur == nil {
            return false
        }
    }

    return cur.isEnd
}


func (this *Trie) StartsWith(prefix string) bool {
    cur := &this.TrieNode
    for _, c := range prefix {
        cur = cur.children[c - 'a']
        if cur == nil {
            return false
        }
    }

    return true
}

