#!/usr/bin/env python3
class LRUCache:
    class ListNode:
        def __init__(self, val=0, key=0, prev=None, next=None):
            self.val = val
            self.key = key
            self.prev = prev
            self.next = next

    def __init__(self, capacity: int):
        self.pointer_cache = {}
        self.capacity = capacity
        self.size = 0
        self.head = None
        self.tail = None

    #head is most recnetly used
    #tail least recently used

    def get(self, key: int) -> int:
        if key in self.pointer_cache:
            node = self.pointer_cache[key]
            if self.size > 1:
                if node == self.tail:
                    self.tail = node.prev
                    self.tail.next = None
                self.head.prev = node
                node.next = self.head
                self.head = node
                node.prev = None

            return node.val

        return -1

    def put(self, key: int, value: int) -> None:
        node = None
        if self.head:
            if key in self.pointer_cache:
                node = self.pointer_cache[key]
                node.val = value
            else:
                node = self.__class__.ListNode(value, key)
                self.pointer_cache[key] = node
                self.size += 1
                if self.size > self.capacity:
                    if key == self.tail.key:
                        self.pointer_cache[key] = node
                    else:
                        del self.pointer_cache[self.tail.key]
                    if self.tail.prev:
                        self.tail = self.tail.prev
                        self.tail.next = None
                    self.size -= 1


            self.head.prev = node
            node.next = self.head
            self.head = node
            node.prev = None

        else:
            node = self.__class__.ListNode(value, key)
            self.pointer_cache[key] = node
            self.head = node
            self.tail = node
            self.head.next = self.tail
            self.tail.prev = self.head
            self.head.prev = None
            self.tail.next = None
            self.size = 1

    def vals(self):
        node = self.head
        while node:
            print(f'{node.key}:{node.val}', end=' ')
            node = node.next
        print()

    def keys(self):
        for key in self.pointer_cache:
            print(f'k-{key}:{self.pointer_cache[key].val}', end=' ')
        print()

cache = LRUCache(2)
cache.put(2, 1)
cache.vals()
cache.put(1, 1)
cache.vals()
cache.put(2, 3)
cache.vals()
cache.keys()
cache.put(4, 1)
cache.vals()
print(cache.get(1))
print(cache.get(2))
cache.keys()